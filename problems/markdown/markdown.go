package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ul   = regexp.MustCompile("(?s)(<li>.*</li>)")
	li   = regexp.MustCompile(`(?m)^\* (.*)\n?$`)
	em   = regexp.MustCompile(`_([^\n]+?)_`)
	bold = regexp.MustCompile(`__([^\n]+?)__`)
)

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = bold.ReplaceAllString(markdown, tag("strong", "$1"))
	markdown = em.ReplaceAllString(markdown, tag("em", "$1"))
	markdown = li.ReplaceAllString(markdown, tag("li", "$1"))
	markdown = ul.ReplaceAllString(markdown, tag("ul", "$1"))
	if markdown[0] == '#' {
		markdown = header(markdown)
	}
	markdown = strings.ReplaceAll(markdown, "\n", "")
	if m, _ := regexp.MatchString("<[hlu]", markdown); m {
		return markdown
	}
	return "<p>" + markdown + "</p>"
}

// tag creates a HTML tag with the name and innerText: e.g. <a>content</a>
func tag(name, content string) string {
	return fmt.Sprintf("<%s>%s</%s>", name, content, name)
}

// header formats MD heading into HTML heading tag.
func header(md string) string {
	lvl := strings.LastIndex(md, "# ") + 1
	nl := strings.Index(md, "\n") - 1
	md = strings.TrimLeft(md, "# ")
	if nl == -2 {
		return fmt.Sprintf("<h%d>%s</h%d>", lvl, md, lvl)
	}
	return fmt.Sprintf("<h%d>%s</h%d>%s", lvl, md[:nl], lvl, md[nl:])
}

// import (
// 	"bufio"
// 	"fmt"
// 	"strings"
// )
// // Render translates markdown to HTML
// func Render(markdown string) string {
// 	markdown = strings.Replace(markdown, "__", "<strong>", 1)
// 	markdown = strings.Replace(markdown, "__", "</strong>", 1)
// 	markdown = strings.Replace(markdown, "_", "<em>", 1)
// 	markdown = strings.Replace(markdown, "_", "</em>", 1)

// 	list := false
// 	html := ""
// 	scanner := bufio.NewScanner(strings.NewReader(markdown))
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		switch {
// 		case line[0] == '#':
// 			html += convertHeaderNeu(line)
// 		case line[0] == '*':
// 			if !list {
// 				html += "<ul>"
// 				list = true
// 			}
// 			html += convertListItemNeu(line)
// 		default:
// 			if list {
// 				html += "</ul>"
// 				list = false
// 			}
// 			html += "<p>" + line + "</p>"
// 		}
// 	}
// 	if list {
// 		html += "</ul>"
// 		list = false
// 	}
// 	return html
// }
// // Converts markdown header that starts with '#' at begin of line to
// // HTML-header.
// func convertHeaderNeu(markdown string) string {
// 	var i int
// 	for i=0; markdown[i] == '#'; i++ { }
// 	return fmt.Sprintf("<h%d>%s</h%d>", i, markdown[i+1:], i)
// }
// // Converts markdown list that starts with '*' at begin of line to
// // HTML-list item.
// func convertListItemNeu(markdown string) string {
// 	if len(markdown) > 1 {
// 		return "<li>" + markdown[2:] + "</li>"
// 	}
// 	return "<li></li>"
// }
