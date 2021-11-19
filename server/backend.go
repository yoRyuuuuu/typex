package server

import (
	"fmt"
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

type AttackAction struct {
	ID     uuid.UUID
	Text   string
	Target string
}

type PlayerInfo struct {
	Name   string
	Health int
}

type Game struct {
	Problem       map[uuid.UUID]IIterator
	PlayerInfo    map[uuid.UUID]*PlayerInfo
	PlayerID      []uuid.UUID
	ActionChannel chan Action
	EventChannel  chan Event
	HasStarted    bool
	PlayerCount   int
	Mu            sync.RWMutex
}

func NewGame() *Game {
	game := &Game{
		Problem:       make(map[uuid.UUID]IIterator),
		PlayerInfo:    make(map[uuid.UUID]*PlayerInfo),
		PlayerID:      []uuid.UUID{},
		ActionChannel: make(chan Action, 1),
		EventChannel:  make(chan Event, 1),
		HasStarted:    false,
		PlayerCount:   0,
		Mu:            sync.RWMutex{},
	}

	return game
}

func (g *Game) Start() {
	go g.watchPlayerCount()
	go g.watchAction()
	go g.watchWinner()
}

func (g *Game) watchPlayerCount() {
	for g.PlayerCount < PlayerCount {
		continue
	}

	time.Sleep(1 * time.Second)
	g.EventChannel <- StartEvent{}
	time.Sleep(5 * time.Second)
	g.HasStarted = true
	for _, id := range g.PlayerID {
		g.Question(id)
	}
}

func (g *Game) watchAction() {
	for {
		action := <-g.ActionChannel
		if !g.HasStarted {
			continue
		}
		g.Mu.Lock()
		action.Perform(g)
		g.Mu.Unlock()
	}
}

func (g *Game) watchWinner() {
	for {
		if !g.HasStarted {
			continue
		}

		g.Mu.RLock()
		// 体力が1以上のプレイヤーが1人のとき終了
		var count = 0
		for _, player := range g.PlayerInfo {
			if player.Health >= 1 {
				count += 1
			}
		}

		if count == 1 {
			for k, player := range g.PlayerInfo {
				if player.Health >= 1 {
					g.EventChannel <- FinishEvent{
						Winner: g.PlayerInfo[k].Name,
					}
					break
				}
			}
		}

		g.Mu.RUnlock()
	}
}

func (game *Game) DamagePlayer(target string) {
	id, _ := uuid.Parse(target)
	game.PlayerInfo[id].Health--
	game.EventChannel <- DamageEvent{
		ID:     target,
		Damage: game.PlayerInfo[id].Health,
	}
}

func (action AttackAction) Perform(game *Game) {
	fmt.Println(action)
	// 不正解ならreturn
	if action.Text != game.Problem[action.ID].Peek() {
		return
	}

	game.Problem[action.ID].Next()
	// ダメージ処理
	game.DamagePlayer(action.Target)
	id, _ := uuid.Parse(action.Target)
	name := game.PlayerInfo[id].Name
	log.Printf("%v's health is %v", name, game.PlayerInfo[action.ID].Health)
	game.Question(action.ID)
}

func (g *Game) Question(id uuid.UUID) {
	g.EventChannel <- QuestionEvent{
		ID:   id,
		Text: g.Problem[id].Peek(),
	}
}
