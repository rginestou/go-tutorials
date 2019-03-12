package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/bradfitz/slice"
	strip "github.com/grokify/html-strip-tags-go"
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
	ID       int
}

type article struct {
	ID      int
	Title   string
	URL     string
	Image   string
	Content string
}

type index struct {
	Articles []article
}

const (
	writeContent int = iota
	writeGuide
	writeTitle
	writeImage
	writeId
)

func main() {
	t, _ := template.ParseFiles("template.html")

	d, _ := os.Open("src/")
	files, _ := d.Readdir(-1)
	d.Close()

	articles := make([]article, 0)

	for _, file := range files {
		if file.Mode().IsRegular() {
			name := file.Name()
			if filepath.Ext(name) == ".md" {
				fmt.Println("Processing", name)

				var l layout
				sections := make([]section, 0)
				s := -1
				writting := writeContent
				id := ""

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
						if comment == "id" {
							writting = writeId
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
					if writting == writeId {
						id += txt
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
				l.ID, _ = strconv.Atoi(id)

				base := name[0 : len(name)-3]
				f, _ := os.Create("www/" + base + ".html")
				w := bufio.NewWriter(f)
				t.Execute(w, l)
				w.Flush()
				f.Close()

				re, _ = regexp.Compile(`<p>(.*?)</p>`)

				ellispsis := ""
				if len(l.Sections) > 0 {
					txt := l.Sections[0].Content
					res := re.FindAllStringSubmatch(txt, -1)
					if len(res) != 0 {
						ellispsis = strip.StripTags(strings.TrimSpace(res[0][1]))
					}
				}
				ll := 180
				if len(ellispsis) < ll {
					ll = len(ellispsis)
				}
				articles = append(articles, article{
					ID:      l.ID,
					Title:   l.Title,
					URL:     base + ".html",
					Image:   l.Image,
					Content: ellispsis[0:ll] + "...",
				})
			}
		}
	}

	slice.Sort(articles[:], func(i, j int) bool {
		return articles[i].ID < articles[j].ID
	})

	// Create index
	l := index{Articles: articles}
	t, _ = template.ParseFiles("index_template.html")
	f, _ := os.Create("www/index.html")
	w := bufio.NewWriter(f)
	t.Execute(w, l)
	w.Flush()
	f.Close()
}
