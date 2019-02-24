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

				sections := make([]section, 0)
				s := -1
				writeContent := true

				c, _ := ioutil.ReadFile(name)
				content := string(blackfriday.Run(c))

				scanner := bufio.NewScanner(strings.NewReader(content))
				for scanner.Scan() {
					txt := scanner.Text()
					if txt == "<p>=== CONTENT ===</p>" {
						sections = append(sections, section{})
						s++
						writeContent = true
						continue
					}
					if txt == "<p>=== GUIDE ===</p>" {
						writeContent = false
						continue
					}

					if writeContent {
						sections[s].Content += txt + "\n"
					} else {
						sections[s].Guide += txt + "\n"
					}
				}

				base := name[0 : len(name)-3]
				f, _ := os.Create("www/" + base + ".html")
				w := bufio.NewWriter(f)
				t.Execute(w, sections)
				w.Flush()
				f.Close()
			}
		}
	}
}
