package backend

type Mode interface{}

type Random struct {
	Mode
}

type Aim struct {
	Mode
	Target int
}
