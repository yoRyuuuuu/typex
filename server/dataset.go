package server

import "math/rand"

type Dataset struct {
	text []string
	idx  int
}

func NewDataset() *Dataset {
	text := []string{
		"apple",
		"orange",
		"stub",
		"mock",
		"golang",
	}

	dataset := &Dataset{
		text: text,
		idx:  rand.Intn(len(text)),
	}

	return dataset
}

func (d *Dataset) Peek() string {
	return d.text[d.idx]
}

func (d *Dataset) Next() string {
	res := d.text[d.idx]
	d.idx = rand.Intn(len(d.text))
	return res
}
