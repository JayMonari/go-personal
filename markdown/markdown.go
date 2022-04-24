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
