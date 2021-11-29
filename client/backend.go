package client

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PlayerInfo struct {
	ID     string
	Name   string
	Health int
}

type Game struct {
	PlayerInfo     map[string]*PlayerInfo
	PlayerID       []string
	Target         string
	EventChannel   chan Event  // GameClientから送信されるEventChannel
	ActionReceiver chan Action // GameClientへ送信するEventChannel
	ActionSender   chan Action
	Problem        string
	Log            string
	ID             string
	Logger         Logger
	Mutex          sync.Mutex
}

func NewGame() *Game {
	game := &Game{
		PlayerInfo:     make(map[string]*PlayerInfo),
		PlayerID:       []string{},
		Target:         "",
		EventChannel:   make(chan Event),
		ActionReceiver: make(chan Action),
		ActionSender:   make(chan Action),
		Logger:         *NewLogger(),
		Problem:        "",
		ID:             "",
		Mutex:          sync.Mutex{},
	}

	go game.watchEvent()
	go game.watchAction()

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
		case DamageEvent:
			g.handleDamageEvent(event)
		}
	}
}

func (g *Game) watchAction() {
	for {
		action := <-g.ActionReceiver
		switch action := action.(type) {
		case Attack:
			g.ActionSender <- Attack{
				Text: action.Text,
				ID:   g.Target,
			}
		case ModeChange:
			g.handleModeChangeAction(action)
		}
	}
}

func (g *Game) handleModeChangeAction(action ModeChange) {
	switch mode := action.Mode.(type) {
	case Random:
		idx := rand.Intn(len(g.PlayerID))
		g.Target = g.PlayerID[idx]
	case Aim:
		if mode.Target < len(g.PlayerID) {
			g.Target = g.PlayerID[mode.Target]
		}
	}
}

func (g *Game) handleDamageEvent(event DamageEvent) {
	g.PlayerInfo[event.ID].Health = event.Damage
}

func (g *Game) handleJoinEvent(event JoinEvent) {
	playerInfo := PlayerInfo{
		ID:     event.ID,
		Name:   event.Name,
		Health: event.Health,
	}
	g.PlayerID = append(g.PlayerID, playerInfo.ID)
	g.PlayerInfo[event.ID] = &playerInfo
}

func (g *Game) handleStartEvent(event StartEvent) {
	idx := rand.Intn(len(g.PlayerID))
	g.Target = g.PlayerID[idx]
	limit := 5 * time.Second
	count := 0
	output := []string{"4", "3", "2", "1", "start!!"}
	for begin := time.Now(); time.Since(begin) < limit; {
		g.Logger.PutString(fmt.Sprintln(output[count]))
		count += 1
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) handleFinishEvent(event FinishEvent) {
	g.Logger.PutString(fmt.Sprintf("Finish! %v Win!!\n", event.Winner))
}

func (g *Game) handleQuestionEvent(event QuestionEvent) {
	g.Problem = event.Text
}

func (g *Game) CheckAnswer(input string) bool {
	return input == g.Problem
}
