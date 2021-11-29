package client

type Action interface{}

type Attack struct {
	Action
	Text string
	ID   string
}

type Mode interface{}

type ModeChange struct {
	Action
	Mode
}

type Random struct {
	Mode
}

type Aim struct {
	Mode
	Target int
}
