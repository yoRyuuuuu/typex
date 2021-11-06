package client

import (
	"bufio"
	"fmt"
	"os"
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
}

type Game struct {
	problem string
	// GameClientから送信されるEvent
	eventChannel chan Event
	// ViewControllerから送信されるAction
	actionChannel chan Action
}

func NewGame() *Game {
	game := &Game{
		problem:       "",
		eventChannel:  make(chan Event),
		actionChannel: make(chan Action),
	}

	go game.watchEvent()
	go game.watchAction()

	return game
}

func (g *Game) Start() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	for {
		if sc.Scan() {
			input := sc.Text()
			if g.checkAnswer(input) {
				g.actionChannel <- AnswerAction{}
			}
		}
	}
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

func (g *Game) watchAction() {
	for {
		action := <-g.actionChannel

		switch action.(type) {

		}
	}
}

func (g *Game) handleStartEvent(event StartEvent) {
	limit := 5 * time.Second
	count := 0
	output := []string{"4", "3", "2", "1", "start!!"}
	for begin := time.Now(); time.Since(begin) < limit; {
		fmt.Println(output[count])
		count += 1
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) handleFinishEvent(event FinishEvent) {
	fmt.Printf("Finish! %v Win!!\n", event.winner)
}

func (g *Game) handleQuestionEvent(event QuestionEvent) {
	g.problem = event.text
	fmt.Printf("%v\n", g.problem)
}

func (g *Game) checkAnswer(input string) bool {
	return input == g.problem
}
