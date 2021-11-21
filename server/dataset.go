package server

import (
	"math/rand"
)

var datasetSingleton *dataset = newDataset()

type dataset struct {
	text []string
}

type IIterator interface {
	Peek() string
	Next() string
}

func newDataset() *dataset {
	return &dataset{
		text: Words,
	}
}

type DatasetIterator struct {
	index int
	dataset
}

func NewDatasetIterator() *DatasetIterator {
	dataset := *datasetSingleton
	index := rand.Intn(len(dataset.text))
	return &DatasetIterator{
		index:   index,
		dataset: *datasetSingleton,
	}
}

func (i *DatasetIterator) Peek() string {
	return i.text[i.index]
}

func (i *DatasetIterator) Next() string {
	text := i.text[i.index]
	i.index = rand.Intn(len(i.text))
	return text
}

func NewDataset() *dataset {
	return datasetSingleton
}
