package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		line = -1
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	w.Header().Set("Content-Type", "text/html")
	formatter := html.New(html.WithLineNumbers(true), html.HighlightLines(lines))
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	iterator, _ := lexers.Get("go").Tokenise(nil, b.String())
	formatter.Format(w, styles.Get("monokai"), iterator)
}

func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func makeLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for i, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := line[1:strings.Index(line, ":")]
		lineNum := strings.LastIndex(line, " ")
		if lineNum == -1 {
			lineNum = len(line)
		}
		lineStr := line[len(file)+2 : lineNum]
		v := url.Values{}
		v.Set("path", file)
		v.Set("line", lineStr)
		lines[i] = fmt.Sprintf(`    <a href="/debug/?%s">%s:%s</a>%s`,
			v.Encode(),    // href...%s
			file, lineStr, // >%s:%s<
			line[len(file)+2+len(lineStr):]) // </a>%s
	}
	return strings.Join(lines, "\n")
}
