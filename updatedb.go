package main

import (

	"os"
	_"path/filepath"
	"strings"
	//"encoding/gob"
	//"bytes"
	"encoding/gob"
	//"path"
	"sync"
	_"sort"

	"path/filepath"

	"fmt"
	"sort"
)


var documents []string


type IndexBuilder interface {
	Build(dir string)
}
type BuildIndexConcurrent struct {}
type BuildIndexWithWalk struct {}

//Builds the index and saves it.
func BuildIndex(dir string){
	config := Config()
	builder := makeIndexBuilder(config)
	builder.Build("/")
	sort.Strings(documents)
	save(documents, config)
}

//makeIndexBuilder is our index builder factory. It chooses the appropriate struct based on the configuration value.
func makeIndexBuilder(c Configuration) IndexBuilder {
	if c.BuildIndexStrategy == "Concurrent" {
		return BuildIndexConcurrent{}
	}
	if c.BuildIndexStrategy == "Iterative" {
		return BuildIndexWithWalk{}
	}
	fmt.Println(os.Stderr, "Invalid configuration value for GOCATE_BUILD_INDEX_STRATEGY. Please set it to Concurrent or Iterative. Choosing Default.")
	return BuildIndexConcurrent{}
}

//Builds index using filepath.Walk.
func (BuildIndexConcurrent) Build(dir string) {
	buildIndexWithWalk(dir)
}

//Builds index using concurrent strategy.
func (BuildIndexWithWalk) Build(dir string) {
	buildIndexConcurrent(dir)
}

//BuildIndexConcurrent builds the index by walking through the filepath while calling readDir concurrently.
//The amount of concurrent calls to readDir are limited for two reasons 1) To prevent opening too many files
//2) To modify for best performance.It is considerably faster than using filepath.Walk when reading large directories.
//Indexing the root directory after wiping filesystem cache (using sync && echo 3 > sudo /proc/sys/vm/drop_caches)
//ran in 18seconds vs 34 seconds with filepath.Walk.
// Using Golangs benchmarking test directly after wiping the filesystem cache yielded:
//BenchmarkBuildIndexWalk-8   	    2000	    508818 ns/op
//BenchmarkBuildIndexConcurrent-8   2000000000	        0.12 ns/op
//Note that these are from different runs to prevent them from interfering which eachother,
func buildIndexConcurrent(dir string) {
	numbJobs := 100
	jobsRunning := 0

	dirsToRead := make(chan string)
	results := make(chan string)
	finished := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)
	jobsRunning += 1

	var dirQueue []string;
	go readDir(dir, results, dirsToRead, finished)

	go func() {
		for {
			select {
			case dir := <-results:
				documents = append(documents, dir)
			case dir := <-dirsToRead:
				dirQueue = append(dirQueue, dir)
			case <-finished:
				jobsRunning--
				processes := min(numbJobs-jobsRunning, len(dirQueue))
				wg.Add(processes)
				for i := 0; i < processes; i++ {
					dir := dirQueue[len(dirQueue) - 1]
					dirQueue = dirQueue[:len(dirQueue) - 1]
					jobsRunning++;
					go readDir(dir, results, dirsToRead, finished);
				}
				wg.Done()
			}
		}
	}()
	wg.Wait()
}


//readDir reads a directory by opening it and iterating over it's files. If the file is a folder
//it is returned on the dirsToRead to be added to the queue of work to do. The function returns all
//paths on the results channel and reports to the finished channel when it is done working.
func readDir(dir string, results chan <- string, dirsToRead chan <- string, finished chan <- bool) {

	file, err := os.Open(dir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		finished <- true
		file.Close()
		return
	}

	slice, err := file.Readdir(-1)

	file.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		finished <- true
		return
	}

	dir = strings.TrimRight(dir, "/")
	for _, fileOrFolder := range slice {
		name := dir + "/" + fileOrFolder.Name()
		if (fileOrFolder.IsDir()) {
			dirsToRead <- name
		}
		results <- name
	}

	finished <- true
}


//BuildIndexWithWalk builds the index using golang's filepath.Walk
func buildIndexWithWalk(dir string) {
	//fmt.Println(len(documents))
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if (err != nil) {
			fmt.Fprintln(os.Stdout, err)
		}
		documents = append(documents, path)
		return nil
	});
}

// Decode Gob filearg
func Load(object interface{}, c Configuration) error {
	file, err := os.Open(c.StorageLocation + "/index.gob")
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

//Encode our index to file
func save(object interface{}, c Configuration) error {
	file, err := os.Create(c.StorageLocation + "/index.gob")
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	} else {
		panic(err)
	}

	file.Close()
	return err
}

func min(x, y int) int {
	if x < y {
		return x

	}
	return y
}
