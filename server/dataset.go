package server

import (
	"bufio"
	"log"
	"math/rand"
	"os"
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
	fp, err := os.Open("./assets/dataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	var text []string
	for sc.Scan() {
		text = append(text, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return &dataset{
		text: text,
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
