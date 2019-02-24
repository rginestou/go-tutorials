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
	Image    string
}

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
				writeContent := true
				writeImage := false

				c, _ := ioutil.ReadFile(name)
				scanner := bufio.NewScanner(strings.NewReader(string(c)))
				for scanner.Scan() {
					txt := scanner.Text()
					if txt == "=== CONTENT ===" {
						sections = append(sections, section{})
						s++
						writeContent = true
						writeImage = false
						continue
					}
					if txt == "=== GUIDE ===" {
						writeContent = false
						writeImage = false
						continue
					}
					if txt == "=== IMAGE ===" {
						writeImage = true
						continue
					}

					if writeImage {
						l.Image += txt
						continue
					}
					if writeContent {
						sections[s].Content += txt + "\n"
					} else {
						sections[s].Guide += txt + "\n"
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
