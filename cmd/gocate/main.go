package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/vlaurenzano/gocate/pkg"
)

func main(){

	if len(os.Args) < 2 {
		fmt.Println("You must supply an argument. Use -h to list instructions.")
		os.Exit(1)
	}



	//all  := flag.Bool("A", false, "Print only entries that match all PATTERNs instead of requiring only one of them to match." );
	update  := flag.Bool("u", false, "Update the database." );

	flag.Parse();

	if *update {
		pkg.BuildIndex("/");
		os.Exit(0);
	}


	var documents []string


	err := pkg.Load(&documents, pkg.Config())

	if(err != nil){
		fmt.Println("There was a problem loading the database, run gocate -u to rebuild it.")
		os.Exit(1)
	}


	results, bestMatches := pkg.MatchFiles( os.Args[1], documents)
	fmt.Println(results)

	printResults(results, bestMatches)

}

func printResults(results []string, matches *pkg.PriorityQueue){

	for _, match := range results {
		fmt.Println(match)
	}

	if !(matches.Len() > 0) {
		return
	}

	fmt.Println("--------------")

	for matches.Len() > 0 {
		item := matches.Pop()
		m := item.(*pkg.Match)
		fmt.Println(m.Value())
	}
}




