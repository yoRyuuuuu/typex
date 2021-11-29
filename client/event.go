package client

type Event interface{}

type FinishEvent struct {
	Event
	Winner string
}

type QuestionEvent struct {
	Event
	Text string
}

type StartEvent struct {
	Event
}

type DamageEvent struct {
	Event
	// 攻撃を受けるPlayerのID
	ID string
	// ダメージの数値
	Damage int
}

type JoinEvent struct {
	Event
	ID     string
	Name   string
	Health int
}
