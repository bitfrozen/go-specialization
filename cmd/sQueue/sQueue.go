package main

import (
	"fmt"
)

type Item struct {
	name string
}

type ItemQueue []*Item

func (q *ItemQueue) top() (topItem *Item) {
	topItem = (*q)[0]
	(*q)[0] = nil
	*q = (*q)[1:]
	return
}

func (q *ItemQueue) push(item *Item) {
	*q = append(*q, item)
}

func (q *ItemQueue) print() {
	for _, v := range *q {
		fmt.Printf("%s:%p,", v.name, v)
	}
	fmt.Println()
}

func main() {
	// create collection
	var collection ItemQueue
	for i := 0; i < 5; i++ {
		collection = append(collection, &Item{name: string(rune(65 + i))})
	}
	collection.print()

	// modify collection
	for i := 0; i < 12; i++ {
		item := collection.top()
		collection.push(item)
	}
	collection.print()

}
