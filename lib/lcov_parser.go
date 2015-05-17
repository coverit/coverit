package main

import (
	"fmt"
)

type Stats struct {
	Lines   int
	Hit     int
	FnFound int
	FnHit   int
	BrFound int
	BrHit   int
}

type Function struct {
	Name string
	Line int
	Hit  int
}

type Line struct {
	Line int
	Hit  int
}

type Branch struct {
	Line   int
	Token  int
	Branch int
	Hit    int
}

type Record struct {
	Name      string
	Test      string
	Stats     Stats
	Functions []Function
	Lines     []Line
	Branches  []Branch
}

type Report []Record

func Parse(lcovFile string) (report Report) {
	fmt.Println(lcovFile)
	return nil
}

func parseRecord(data string) (record Record) {
	fmt.Println(data)
	return Record{}
}
