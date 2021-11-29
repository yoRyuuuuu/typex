package client

type Event interface{}

// ゲーム終了Event
type FinishEvent struct {
	Event
	Winner string
}

// お題Event
type QuestionEvent struct {
	Event
	Text string
}

// ゲーム開始Event
type StartEvent struct {
	Event
}

// ダメージEvent
type DamageEvent struct {
	Event
	// 攻撃を受けるPlayerID
	ID string
	// ダメージの数値
	Damage int
}

// プレイヤー参加Event
type JoinEvent struct {
	Event
	// PlayerID
	ID string
	// Playerの名前
	Name string
	// 初期体力
	Health int
}
