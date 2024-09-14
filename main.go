package main

import (
	"fmt"
	"os"
	"flag"
	"bytes"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<meta http-equiv="content-type" content="text/html; charset=utf-8">
<title>Markdown Preview Tool</title>
</head>
<body>`
	footer = "</body></html>"
)

func main(){
	filename := flag.String("file", "", "Name of the markdown file")
	flag.Parse()

	if *filename == ""{
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Println("Error processing the file")
		os.Exit(1)
	}
}

func run(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		return err
	}
	html := parseContent(file)
	htmlFilename := fmt.Sprintf("%s.html", filepath.Base(filename))
	saveHTML(htmlFilename, html)
	return nil
}

func parseContent(file []byte) []byte{
	output := blackfriday.Run(file)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	var buffer bytes.Buffer

	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(filename string, body []byte){
	os.WriteFile(filename, body, 0644)
}

