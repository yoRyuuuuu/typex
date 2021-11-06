package client

import (
	"fmt"
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

type Game struct {
	word string
	// GameClientから送信されるEvent
	eventChannel  chan Event
	actionChannel chan Action
	log           string
}

func NewGame() *Game {
	game := &Game{
		word:          "",
		eventChannel:  make(chan Event),
		actionChannel: make(chan Action),
	}

	go game.watchEvent()

	return game
}

func (g *Game) watchEvent() {
	for {
		event := <-g.eventChannel

		switch event.(type) {
		case FinishEvent:
			event := event.(FinishEvent)
			g.handleFinishEvent(event)
		case QuestionEvent:
			event := event.(QuestionEvent)
			g.handleQuestionEvent(event)
		case StartEvent:
			event := event.(StartEvent)
			g.handleStartEvent(event)
		}
	}
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
