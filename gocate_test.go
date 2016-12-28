package main

import "testing"
//import "reflect"


//func TestMainFunc(t *testing.T) {
//
//}




func BenchmarkReadDirWithWak(b *testing.B) {
	buildIndexWithWalk("/home/vlaurenzano")
}

func BenchmarkBuildIndex(b *testing.B) {
	buildIndex("/")
}

//func BenchmarkReadDirWithReadDir(b *testing.B){
//	go readDirWithReadDir("/home/vlaurenzano")
//}
//
//
//func TestReadDirWithReadDir(b *testing.T){
//	readDirWithReadDir("/home/vlaurenzano")
//}


func TestBuildIndex(b *testing.T){
	buildIndex("/home/vlaurenzano")
}
