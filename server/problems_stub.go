package server

import "math/rand"

type Problems interface {
	Peek() string
	Next() string
}

type StubProblems struct {
	text []string
	idx  int
}

func NewStubProblems() *StubProblems {
	text := []string{
		"apple",
		"orange",
		"stub",
		"mock",
		"golang",
	}

	problems := &StubProblems{
		text: text,
		idx:  rand.Intn(len(text)),
	}

	return problems
}

func (p *StubProblems) Peek() string {
	return p.text[p.idx]
}

func (p *StubProblems) Next() string {
	res := p.text[p.idx]
	p.idx = rand.Intn(len(p.text))
	return res
}
