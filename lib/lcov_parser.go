package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
	Taken  int
	Branch int
	Block  int
}

type Record struct {
	Name      string
	File      string
	Stats     Stats
	Functions []Function
	Lines     []Line
	Branches  []Branch
}

type Report []Record

func parseRecord(data string) (record Record) {
	fmt.Println(data)
	data = strings.TrimSpace(data)
	rows := strings.Split(data, "\n")

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		method, content := tupleFromString(row, ":")

		if method == "TN" {
			// test title
			record.Name = content
		} else if method == "SF" {
			// file name
			record.File = content
		} else if method == "LF" {
			// line found
			if lf, err := strconv.Atoi(content); err == nil {
				record.Stats.Lines = lf
			}
		} else if method == "LH" {
			// hit found
			if lh, err := strconv.Atoi(content); err == nil {
				record.Stats.Hit = lh
			}
		} else if method == "DA" {
			line, hit := tupleFromString(content, ",")
			var lineData Line
			if lineValue, err := strconv.Atoi(line); err == nil {
				lineData.Line = lineValue
			}
			if hitValue, err := strconv.Atoi(hit); err == nil {
				lineData.Hit = hitValue
			}
			record.Lines = append(record.Lines, lineData)
		} else if method == "FNF" {
			// ---------
			// Functions
			// ---------
			// functions found
			if ff, err := strconv.Atoi(content); err == nil {
				record.Stats.FnFound = ff
			}
		} else if method == "FNH" {
			// functions hit
			if fh, err := strconv.Atoi(content); err == nil {
				record.Stats.FnHit = fh
			}
		} else if method == "FN" {
			// functions
			line, name := tupleFromString(content, ",")
			var fnData Function
			if lineValue, err := strconv.Atoi(line); err == nil {
				fnData.Line = lineValue
			}
			fnData.Name = name
			record.Functions = append(record.Functions, fnData)
		} else if method == "FNDA" {
			// functions data
			hit, name := tupleFromString(content, ",")
			if hitValue, err := strconv.Atoi(hit); err == nil {
				for i := 0; i < len(record.Functions); i++ {
					fn := record.Functions[i]
					if fn.Name == name {
						fn.Hit = hitValue
					}
				}
			}
		} else if method == "BRF" {
			// ---------
			// Branches
			// ---------
			// branches found
			if bf, err := strconv.Atoi(content); err == nil {
				record.Stats.BrFound = bf
			}
		} else if method == "BRH" {
			// branches hit
			if bh, err := strconv.Atoi(content); err == nil {
				record.Stats.BrHit = bh
			}
		} else if method == "BRDA" {
			// branches data
			parts := strings.Split(content, ",")
			line := parts[0]
			block := parts[1]
			branch := parts[2]
			taken := parts[3]

			var brData Branch
			if lineValue, err := strconv.Atoi(line); err == nil {
				brData.Line = lineValue
			}
			if blockValue, err := strconv.Atoi(block); err == nil {
				brData.Block = blockValue
			}
			if branchValue, err := strconv.Atoi(branch); err == nil {
				brData.Branch = branchValue
			}
			if takenValue, err := strconv.Atoi(taken); err == nil {
				brData.Taken = takenValue
			}
			record.Branches = append(record.Branches, brData)
		} else {
			fmt.Println("Invalid method name [%s]", method)
		}
	}

	return
}

func ParseLcov(lcovFile string) (report Report, err error) {
	data, err := ioutil.ReadFile(lcovFile)
	if err != nil {
		panic(err)
	}
	content := strings.TrimSpace(string(data))
	records := strings.Split(content, "end_of_record")
	for i := 0; i < len(records); i++ {
		report = append(report, parseRecord(records[i]))
	}
	fmt.Println("Awesome, check your report = \n%s", report)
	return
}

func tupleFromString(data string, separator string) (head string, tail string) {
	if len(data) == 0 {
		head = ""
		tail = ""
		return
	}
	parts := strings.Split(data, separator)
	head = parts[0]
	tail = parts[1]
	return
}
