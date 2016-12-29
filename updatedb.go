package main

import (
	"fmt"
	"os"
	_"path/filepath"
	"strings"
	//"encoding/gob"
	//"bytes"
	"encoding/gob"
	//"path"
	"sync"
	"sort"
	_"time"
	"path/filepath"
)

var documents []string

const DB_PATH = "/tmp/docs.gob"


var folderQueue = make(chan string, 10)


func addToQueue(path string){
	folderQueue <- path
}


func Build() {
	buildIndexWithWalk("/")
	Save(documents)

}


func buildIndex(dir string) {
	c := make(chan string, 10)

	defer close(c)

	var wg sync.WaitGroup
	wg.Add(1)

	go readDirWithReadDir(dir, c, &wg);
	go func() {
		for path := range c {
			documents = append(documents, path)
			matched := strings.Contains(path, "hello")
			if matched {
			//	fmt.Println(path)

			}

		}
	}()

	wg.Wait()

	sort.Strings(documents)

	for _ , path := range documents {
		fmt.Println(path)
	}

	Save(documents)
}


func buildIndexWithWalk(dir string) {
	//fmt.Println(len(documents))
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if(err != nil){
			fmt.Println(err)
		}
		documents = append(documents, path)
		return nil
	});
}



func readDirWithReadDir(dir string, c chan string,wg *sync.WaitGroup) {

	defer wg.Done()

	file, err := os.Open(dir)


	if err != nil {
		fmt.Println(err)
	}
//buildIndex("/home/vlaurenzano")
	//a, _ := os.Lstat("/home/vlaurenzano")
	slice, err := file.Readdir(-1)

	file.Close()

	if err != nil {
		fmt.Println(err)
	}

	_ = err

	var folders []string

	for _, fileOrFolder := range slice {

		name := dir + "/" + fileOrFolder.Name()

		if(fileOrFolder.IsDir()) {
			folders = append(folders, name)
		} else {
			c<-name
		}

	}

	if l := len(folders); l > 0 {
		wg.Add(l)
		for _, folder := range folders {
			go readDirWithReadDir(folder, c, wg)
		}
	}

}



// Decode Gob filearg
func Load(object interface{}) error {
	file, err := os.Open(DB_PATH)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// Encode via Gob to file
func Save(object interface{}) error {
	file, err := os.Create(DB_PATH)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	} else {
		panic(err)
	}

	file.Close()
	return err
}


