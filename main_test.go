package main

import (
	"fmt"
	"testing"
	"os"
	"bytes"
)

const (
	inputFile = "./testdata/test.md"
	resultFile = "./test.md.html"
	goldenFile = "./testdata/test.md.html"
)

func TestParseContent(t *testing.T){
	file, err := os.ReadFile(inputFile)

	if err != nil {
		t.Fatal(err)
	}

	expected, _ := os.ReadFile(goldenFile)

	result := parseContent(file)

	if !bytes.Equal(expected, result){
		fmt.Println("error")
		fmt.Println(expected, len(expected))
		fmt.Println(result, len(result))
		t.Error("Result file is not correct")
	}
}

func TestRun(t *testing.T){
	err := run(inputFile)	

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(resultFile)

	file, _ := os.ReadFile(resultFile)

	expected, _ := os.ReadFile(goldenFile)

	if !bytes.Equal(expected, file){
		t.Error("Files are not equal")
	}
}
