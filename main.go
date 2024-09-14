package main

import (
	"io"
	"fmt"
	"os"
	"flag"
	"bytes"

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

	if err := run(*filename, os.Stdout); err != nil {
		fmt.Println("Error processing the file")
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		return err
	}
	html := parseContent(file)

	temp, err := os.CreateTemp("", "md-*.html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	defer temp.Close()

	htmlFileName := temp.Name()

	fmt.Fprintln(out, htmlFileName)

	saveHTML(htmlFileName, html)
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

