package main

import (
	"strings"
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
	var buf bytes.Buffer
	err := run(inputFile, &buf)	

	if err != nil {
		t.Fatal(err)
	}

	tempFile := strings.TrimSpace(buf.String())
	defer os.Remove(tempFile)

	file, err := os.ReadFile(tempFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	expected, _ := os.ReadFile(goldenFile)

	if !bytes.Equal(expected, file){
		fmt.Printf("%s", expected)
		fmt.Println("---------")
		fmt.Printf("%s", file)
		t.Error("Files are not equal")
	}
}
