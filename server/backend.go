package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

const MaxPlayer = 2

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Action interface {
	Perform(game *Game)
}

type Event interface{}

// 次の問題を返す
// 得点を加算する
// 勝者を決める
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
}

func (g *Game) watchPlayerCount() {
	for g.Player < MaxPlayer {
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

func (game *Game) AddScore(id uuid.UUID) {
	game.Score[id]++
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
