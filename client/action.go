package client

type Action interface{}

type Attack struct {
	Action
	Text string
}

type ModeChange struct {
	Action
	Mode
}

type Mode interface{}

type Random struct {
	Mode
}

type Aim struct {
	Mode
	Target int
}
