package main

import (
	"testing"
)


//Benchmark Building and Index with the Walk Function
func BenchmarkBuildIndexWalk(b *testing.B) {
	buildIndexWithWalk("/home")
}

func BenchmarkBuildIndexConcurrent(b *testing.B) {
	buildIndexConcurrent("/home")
}

func TestBuildIndexConcurrent(b *testing.T) {
	buildIndexConcurrent("/home")
}
func TestBuildIndexWalk(b *testing.T) {
	buildIndexWithWalk("/home")
}
