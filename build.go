package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type section struct {
	Content string
	Guide   string
}

type layout struct {
	Sections []section
	Title    string
	Image    string
}

const (
	writeContent int = iota
	writeGuide
	writeTitle
	writeImage
)

func main() {
	t, _ := template.ParseFiles("template.html")

	d, _ := os.Open("./")
	files, _ := d.Readdir(-1)
	d.Close()

	for _, file := range files {
		if file.Mode().IsRegular() {
			name := file.Name()
			if filepath.Ext(name) == ".md" {
				fmt.Println("Processing", name)

				var l layout
				sections := make([]section, 0)
				s := -1
				writting := writeContent

				c, _ := ioutil.ReadFile(name)
				scanner := bufio.NewScanner(strings.NewReader(string(c)))
				for scanner.Scan() {
					txt := scanner.Text()
					if txt == "=== CONTENT ===" {
						writting = writeContent
						sections = append(sections, section{})
						s++
						continue
					}
					if txt == "=== GUIDE ===" {
						writting = writeGuide
						continue
					}
					if txt == "=== TITLE ===" {
						writting = writeTitle
						continue
					}
					if txt == "=== IMAGE ===" {
						writting = writeImage
						continue
					}

					if writting == writeContent {
						sections[s].Content += txt + "\n"
						continue
					}
					if writting == writeGuide {
						sections[s].Guide += txt + "\n"
						continue
					}
					if writting == writeTitle {
						l.Title += txt
						continue
					}
					if writting == writeImage {
						l.Image += txt
						continue
					}
				}

				for s := range sections {
					sections[s].Content = string(blackfriday.Run([]byte(sections[s].Content)))
					sections[s].Guide = string(blackfriday.Run([]byte(sections[s].Guide)))
				}

				l.Sections = sections

				base := name[0 : len(name)-3]
				f, _ := os.Create("www/" + base + ".html")
				w := bufio.NewWriter(f)
				t.Execute(w, l)
				w.Flush()
				f.Close()
			}
		}
	}
}
