package main

import (
	"os"

	"fmt"
	"strings"
)

func main(){

	if len(os.Args) < 2 {
		panic("You must supply an argument")
	}

	var documents []string

	err := Load(&documents)
	if(err != nil){
		panic(err)
	}

	pattern := os.Args[1]

	matchFiles(pattern, documents)
}



func matchFiles(pattern string, documents []string) {

	for _, path := range documents {
		matched := strings.Contains(path, pattern)
		if matched {
			fmt.Println(path)
		}
	}

}