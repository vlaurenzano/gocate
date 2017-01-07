package pkg

import (
	"strings"
	"container/heap"
)

func MatchFiles(pattern string, documents []string)  (results []string, best *PriorityQueue) {

	matches := new(PriorityQueue)

	for _, path := range documents {
		idx := strings.Index(path, pattern)
		if idx >= 0 {
			results = append(results, path)
			item := Match{value: path, priority:(len(path) - idx)}
			matches.Push(&item)
			if matches.Len() > 5 {
				heap.Init(matches)
				heap.Pop(matches)
			}
		}
	}

	return results, matches

}

