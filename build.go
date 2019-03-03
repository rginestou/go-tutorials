package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

	d, _ := os.Open("src/")
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

				c, _ := ioutil.ReadFile("src/" + name)
				scanner := bufio.NewScanner(strings.NewReader(string(c)))
				re, _ := regexp.Compile(`<!--(.*)-->`)

				for scanner.Scan() {
					txt := scanner.Text()
					res := re.FindAllStringSubmatch(txt, -1)
					if len(res) != 0 {
						comment := strings.TrimSpace(res[0][1])
						if comment == "content" {
							writting = writeContent
							sections = append(sections, section{})
							s++
							continue
						}
						if comment == "guide" {
							writting = writeGuide
							continue
						}
						if comment == "title" {
							writting = writeTitle
							continue
						}
						if comment == "image" {
							writting = writeImage
							continue
						}
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

					sections[s].Content = strings.Replace(sections[s].Content, " :</p>", "&nbsp;:</p>", -1)
					sections[s].Guide = strings.Replace(sections[s].Guide, " :</p>", "&nbsp;:</p>", -1)

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
