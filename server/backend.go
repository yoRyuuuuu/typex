package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

const MaxScore = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Action interface {
	Perform(game *Game)
}

type Event interface{}

type Game struct {
	Score         map[uuid.UUID]int
	Problems      map[uuid.UUID]Problems
	Name          map[uuid.UUID]string
	ActionChannel chan Action
	EventChannel  chan Event
	Mu            sync.RWMutex
	WaitForRound  bool
	Player        int
}

func NewGame() *Game {
	game := &Game{
		Score:         make(map[uuid.UUID]int),
		Problems:      make(map[uuid.UUID]Problems),
		ActionChannel: make(chan Action, 1),
		EventChannel:  make(chan Event, 1),
		Name:          make(map[uuid.UUID]string),
		WaitForRound:  true,
	}

	return game
}

func (g *Game) Start() {
	go g.watchPlayerCount()
	go g.watchAction()
	go g.watchWinner()
}

func (g *Game) watchPlayerCount() {
	for g.Player < maxPlayer {
		continue
	}

	time.Sleep(1 * time.Second)
	g.EventChannel <- StartEvent{}
	time.Sleep(5 * time.Second)
	g.WaitForRound = false
	for id := range g.Name {
		g.Question(id)
	}
}

func (g *Game) watchAction() {
	for {
		action := <-g.ActionChannel
		if g.WaitForRound {
			continue
		}
		g.Mu.Lock()
		action.Perform(g)
		g.Mu.Unlock()
	}
}

func (g *Game) watchWinner() {
	for {
		g.Mu.RLock()
		for k, v := range g.Score {
			if v >= MaxScore {
				g.EventChannel <- FinishEvent{
					Winner: g.Name[k],
				}
				return
			}
		}
		g.Mu.RUnlock()
	}
}

type FinishEvent struct {
	Event
	Winner string
}

type StartEvent struct {
	Event
}

type QuestionEvent struct {
	Event
	ID   uuid.UUID
	Text string
}

type AnswerAction struct {
	ID uuid.UUID
}

type DamageEvent struct {
	Event
	id     string
	damage int
}

func (game *Game) AddScore(id uuid.UUID) {
	game.Score[id]++
	game.EventChannel <- DamageEvent{
		id:     id.String(),
		damage: game.Score[id],
	}
}

func (action AnswerAction) Perform(game *Game) {
	game.AddScore(action.ID)
	name := game.Name[action.ID]
	log.Printf("%v's score is %v", name, game.Score[action.ID])
	game.Question(action.ID)
}

func (g *Game) Question(id uuid.UUID) {
	problems := g.Problems[id]
	g.EventChannel <- QuestionEvent{
		ID:   id,
		Text: problems.Next(),
	}
}
