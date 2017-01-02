package main

import (
	"testing"
	"reflect"
)

//Benchmark Building and Index with the Walk Function
func BenchmarkBuildIndexWalk(b *testing.B) {
	buildIndexWithWalk("/home")
}

func BenchmarkBuildIndexConcurrent(b *testing.B) {
	buildIndexConcurrent("/home")
}

func TestMakeIndexBuilder(t *testing.T) {
	c := Configuration{BuildIndexStrategy: "Concurrent"}
	i := makeIndexBuilder(c)
	if reflect.TypeOf(i).Name() != "BuildIndexConcurrent" {
		t.Error("MakeIndexBuilder returned the wrong Builder. Received: ", reflect.TypeOf(i).Name(), " Expected: BuildIndexConcurrent")

	}
	c = Configuration{BuildIndexStrategy: "Iterative"}
	i = makeIndexBuilder(c)
	if reflect.TypeOf(i).Name() != "BuildIndexWithWalk" {
		t.Error("MakeIndexBuilder returned the wrong Builder. Received: ", reflect.TypeOf(i).Name(), " Expected: BuildIndexWithWalk")

	}
}

func TestBuildIndexConcurrent(b *testing.T) {
	buildIndexConcurrent("/home")
}
func TestBuildIndexWithWalk(b *testing.T) {
	buildIndexWithWalk("/home")
}
