package main

import (
	"fmt"
	"html"
)

func main() {
}

func UnescapeStringExample() {
	fmt.Println(html.EscapeString(`&amp;hearts:&#34;Gooey &amp; Gooper&#39;s Hotel&#34; &lt;gotel@example.net&gt;`))
}

func EscapeStringExample() {
	fmt.Println(html.EscapeString(`"Gooey & Gooper's Hotel" <gotel@example.net>`))
}

func NotEqualExample() {
	fmt.Println(html.EscapeString(html.UnescapeString(`&hearts:"Gooey & Gooper's Hotel" <gotel@example.net>`)), "!=", html.UnescapeString(html.EscapeString(`&hearts:"Gooey & Gooper's Hotel" <gotel@example.net>`)))
}
