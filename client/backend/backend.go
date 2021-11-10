package backend

import (
	"fmt"
	"sync"
	"time"

	. "github.com/yoRyuuuuu/typex/common"
)

type Player struct {
	ID   string
	Name string
}

type Action interface{}

type Answer struct {
	Action
	text string
}

type Game struct {
	// GameClientから送信されるEvent
	Players       map[string]Player
	Health        map[string]int
	EventChannel  chan Event
	ActionChannel chan Action
	Problem       string
	Log           string
	ID            string
	Mutex         sync.Mutex
}

func NewGame() *Game {
	game := &Game{
		Players:       map[string]Player{},
		Health:        map[string]int{},
		EventChannel:  make(chan Event),
		ActionChannel: make(chan Action),
		Problem:       "",
		Log:           "",
		ID:            "",
		Mutex:         sync.Mutex{},
	}

	go game.watchEvent()

	return game
}

func (g *Game) watchEvent() {
	for {
		event := <-g.EventChannel
		switch event := event.(type) {
		case FinishEvent:
			g.handleFinishEvent(event)
		case QuestionEvent:
			g.handleQuestionEvent(event)
		case StartEvent:
			g.handleStartEvent(event)
		case JoinEvent:
			g.handleJoinEvent(event)
		case AttackEvent:
			g.handleAttackEvent(event)
		}
	}
}

func (g *Game) handleAttackEvent(event AttackEvent) {
	g.Health[event.ID] = event.Health
}

func (g *Game) handleJoinEvent(event JoinEvent) {
	player := Player{
		ID:   event.ID,
		Name: event.Name,
	}
	g.Players[event.ID] = player
}

func (g *Game) handleStartEvent(event StartEvent) {
	limit := 5 * time.Second
	count := 0
	output := []string{"4", "3", "2", "1", "start!!"}
	for begin := time.Now(); time.Since(begin) < limit; {
		g.Log = output[count]
		count += 1
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) handleFinishEvent(event FinishEvent) {
	g.Log = fmt.Sprintf("Finish! %v Win!!\n", event.Winner)
}

func (g *Game) handleQuestionEvent(event QuestionEvent) {
	g.Problem = event.Text
}

func (g *Game) CheckAnswer(input string) bool {
	return input == g.Problem
}
