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
	word string
	// GameClientから送信されるEvent
	eventChannel  chan Event
	actionChannel chan Action
	log           string
	players       map[string]Player
	mutex         sync.Mutex
	id            string
	health        map[string]int
}

func NewGame() *Game {
	game := &Game{
		word:          "",
		eventChannel:  make(chan Event),
		actionChannel: make(chan Action),
		log:           "",
		players:       map[string]Player{},
		mutex:         sync.Mutex{},
		id:            "",
		health:        map[string]int{},
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
	g.word = event.Text
}

func (g *Game) checkAnswer(input string) bool {
	return input == g.word
}
