package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

const MaxScore = 10
const InitialHealth = 15

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Action interface {
	Perform(game *Game)
}

type Event interface{}

type Game struct {
	Health        map[uuid.UUID]int
	Problems      map[uuid.UUID]Problems
	Name          map[uuid.UUID]string
	PlayerID      []uuid.UUID
	ActionChannel chan Action
	EventChannel  chan Event
	Mu            sync.RWMutex
	WaitForRound  bool
	PlayerCount   int
}

func NewGame() *Game {
	game := &Game{
		Health:        make(map[uuid.UUID]int),
		Problems:      make(map[uuid.UUID]Problems),
		Name:          make(map[uuid.UUID]string),
		PlayerID:      []uuid.UUID{},
		ActionChannel: make(chan Action, 1),
		EventChannel:  make(chan Event, 1),
		Mu:            sync.RWMutex{},
		WaitForRound:  true,
		PlayerCount:   0,
	}

	return game
}

func (g *Game) Start() {
	go g.watchPlayerCount()
	go g.watchAction()
	go g.watchWinner()
}

func (g *Game) watchPlayerCount() {
	for g.PlayerCount < maxPlayer {
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
		if g.WaitForRound {
			continue
		}

		g.Mu.RLock()
		// 体力が1以上のプレイヤーが1人のとき終了
		var count = 0
		for _, v := range g.Health {
			if v >= 1 {
				count += 1
			}
		}

		if count == 1 {
			for k, v := range g.Health {
				if v >= 1 {
					g.EventChannel <- FinishEvent{
						Winner: g.Name[k],
					}
				}
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

type AttackEvent struct {
	Event
	id     string
	health int
}

func (game *Game) AttackPlayer(id uuid.UUID) {
	game.Health[id]--
	game.EventChannel <- AttackEvent{
		id:     id.String(),
		health: game.Health[id],
	}
}

func (action AnswerAction) Perform(game *Game) {
	idx := rand.Intn(len(game.PlayerID))
	game.AttackPlayer(game.PlayerID[idx])
	name := game.Name[action.ID]
	log.Printf("%v's health is %v", name, game.Health[action.ID])
	game.Question(action.ID)
}

func (g *Game) Question(id uuid.UUID) {
	problems := g.Problems[id]
	g.EventChannel <- QuestionEvent{
		ID:   id,
		Text: problems.Next(),
	}
}
