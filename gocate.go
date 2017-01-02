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



