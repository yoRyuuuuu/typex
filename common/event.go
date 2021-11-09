package common

import "github.com/google/uuid"

type Event interface{}

type FinishEvent struct {
	Event
	Winner string
}

type QuestionEvent struct {
	Event
	ID   uuid.UUID
	Text string
}

type StartEvent struct {
	Event
}

type AttackEvent struct {
	Event
	// 攻撃を受けるPlayerのID
	ID string
	// 体力の数値
	Health int
}

type JoinEvent struct {
	Event
	ID   string
	Name string
}
