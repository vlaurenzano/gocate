package main

import (
	"os"
	"fmt"
	"strings"
	"container/heap"
	"flag"
)

func main(){

	if len(os.Args) < 2 {
		fmt.Println("You must supply an argument. Use -h to list instructions.")
		os.Exit(1)
	}


	update  := flag.Bool("u", false, "Update the database" );

	flag.Parse();

	if *update {
		BuildIndex("/");
		os.Exit(0);
	}




	var documents []string


	err := Load(&documents, Config())

	if(err != nil){
		panic(err)
	}

	pattern := os.Args[1]

	matchFiles(pattern, documents)

}


func matchFiles(pattern string, documents []string) {

	matches := new(PriorityQueue)

	for _, path := range documents {
		idx := strings.Index(path, pattern)
		if idx >= 0 {
			fmt.Println(path)
			item := Match{value: path, priority:(len(path) - idx)}
			matches.Push(&item)

			if matches.Len() > 5 {
				heap.Init(matches)
				heap.Pop(matches)
			}

		}
	}


	fmt.Println("--------------")

	for matches.Len() > 0 {
			item := matches.Pop()
			m := item.(*Match)
			fmt.Println(m.value)
	}

}



// An Item is something we manage in a priority queue.
type Match struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Match

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Match)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Match, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

