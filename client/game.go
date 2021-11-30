package client

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/yoRyuuuuu/typex/proto"
	"google.golang.org/grpc/metadata"
)

type PlayerStatus struct {
	ID     string
	Name   string
	Health int
}

type Game struct {
	ActionReceiver chan Action
	PlayerStatuses map[string]*PlayerStatus
	EnemyIDs       []string
	MyID           string
	Target         string
	Word           string
	Logger         Logger
	Mutex          sync.Mutex
	*GameClient
}

func NewGame(gameClient *GameClient) *Game {
	game := &Game{
		ActionReceiver: make(chan Action),
		PlayerStatuses: make(map[string]*PlayerStatus),
		MyID:           "",
		EnemyIDs:       []string{},
		Target:         "",
		Word:           "",
		Logger:         *NewLogger(),
		Mutex:          sync.Mutex{},
		GameClient:     gameClient,
	}
	return game
}

func (g *Game) Start() {
	go g.watchEvent()
	go g.watchAction()
}

// サーバ接続処理
func (g *Game) Connect(grpcClient proto.GameClient, playerName string) error {
	req := proto.ConnectRequest{
		Name: playerName,
	}
	res, err := grpcClient.Connect(context.Background(), &req)
	if err != nil {
		return err
	}

	g.MyID = res.GetId()
	for _, player := range res.GetPlayer() {
		status := &PlayerStatus{
			ID:     player.Id,
			Name:   player.Name,
			Health: int(player.Health),
		}
		g.PlayerStatuses[player.Id] = status
		if status.ID != g.MyID {
			g.EnemyIDs = append(g.EnemyIDs, status.ID)
		}
	}

	header := metadata.New(map[string]string{"authorization": res.Id})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	stream, err := grpcClient.Stream(ctx)
	if err != nil {
		return err
	}

	g.Stream = stream
	return nil
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
			g.handleAttackAction(action.Text, g.Target)
		case ModeChange:
			g.handleModeChangeAction(action)
		}
	}
}

func (g *Game) handleModeChangeAction(action ModeChange) {
	switch mode := action.Mode.(type) {
	case Random:
		id := g.EnemyIDs[rand.Intn(len(g.EnemyIDs))]
		g.Target = g.PlayerStatuses[id].ID
	case Aim:
		if mode.Target < len(g.EnemyIDs) {
			g.Target = g.EnemyIDs[mode.Target]
		}
	}
}

func (g *Game) handleDamageEvent(event DamageEvent) {
	g.PlayerStatuses[event.ID].Health = event.Damage
}

func (g *Game) handleJoinEvent(event JoinEvent) {
	playerInfo := PlayerStatus{
		ID:     event.ID,
		Name:   event.Name,
		Health: event.Health,
	}
	g.EnemyIDs = append(g.EnemyIDs, playerInfo.ID)
	g.PlayerStatuses[event.ID] = &playerInfo
}

func (g *Game) handleStartEvent(event StartEvent) {
	idx := rand.Intn(len(g.EnemyIDs))
	g.Target = g.EnemyIDs[idx]
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
	g.Logger.PutString(fmt.Sprintln("Press contrl+c to exit"))
}

func (g *Game) handleQuestionEvent(event QuestionEvent) {
	g.Word = event.Text
}
