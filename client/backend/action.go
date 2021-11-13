package backend

type Action interface{}

type Attack struct {
	Action
	Text string
	ID   string
}

type ModeChange struct {
	Action
	Mode
}
