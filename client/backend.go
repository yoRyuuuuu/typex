package client

import (
	"fmt"
	"sync"
	"time"
)

type Event interface{}

type Action interface{}

type FinishEvent struct {
	Event
	winner string
}

type QuestionEvent struct {
	Event
	text string
}

type StartEvent struct {
	Event
}

type AnswerAction struct {
	Action
	text string
}

type JoinEvent struct {
	Event
	id   string
	name string
}

type Player struct {
	id    string
	name  string
}

type DamageEvent struct {
	Event
	id     string
	damage int
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
	score         map[string]int
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
		score:         map[string]int{},
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
		case DamageEvent:
			g.handleDamageEvent(event)
		}
	}
}

func (g *Game) handleDamageEvent(event DamageEvent) {
	g.score[event.id] = event.damage
}

func (g *Game) handleJoinEvent(event JoinEvent) {
	player := Player{
		id:   event.id,
		name: event.name,
	}
	g.players[event.id] = player
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
	g.log = fmt.Sprintf("Finish! %v Win!!\n", event.winner)
}

func (g *Game) handleQuestionEvent(event QuestionEvent) {
	g.word = event.text
}

func (g *Game) checkAnswer(input string) bool {
	return input == g.word
}
