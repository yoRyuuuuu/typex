package client

import (
	"fmt"
	"sync"
	"time"

	. "github.com/yoRyuuuuu/typex/common"
)

type Player struct {
	id   string
	name string
}

type Action interface{}

type AnswerAction struct {
	Action
	text string
}

type Game struct {
	// GameClientから送信されるEvent
	players       map[string]Player
	health        map[string]int
	eventChannel  chan Event
	actionChannel chan Action
	problem          string
	log           string
	id            string
	mutex         sync.Mutex
}

func NewGame() *Game {
	game := &Game{
		players:       map[string]Player{},
		health:        map[string]int{},
		eventChannel:  make(chan Event),
		actionChannel: make(chan Action),
		problem:          "",
		log:           "",
		id:            "",
		mutex:         sync.Mutex{},
	}

	go game.watchEvent()

	return game
}

func (g *Game) watchEvent() {
	for {
		event := <-g.eventChannel
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
	g.health[event.ID] = event.Health
}

func (g *Game) handleJoinEvent(event JoinEvent) {
	player := Player{
		id:   event.ID,
		name: event.Name,
	}
	g.players[event.ID] = player
}

func (g *Game) handleStartEvent(event StartEvent) {
	limit := 5 * time.Second
	count := 0
	output := []string{"4", "3", "2", "1", "start!!"}
	for begin := time.Now(); time.Since(begin) < limit; {
		g.log = output[count]
		count += 1
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) handleFinishEvent(event FinishEvent) {
	g.log = fmt.Sprintf("Finish! %v Win!!\n", event.Winner)
}

func (g *Game) handleQuestionEvent(event QuestionEvent) {
	g.problem = event.Text
}

func (g *Game) checkAnswer(input string) bool {
	return input == g.problem
}
