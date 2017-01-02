package main

import "testing"



func TestMainFunc(t *testing.T) {
	var documents []string
	_ = Load(&documents, Config())
	matchFiles("hello", documents )
}



