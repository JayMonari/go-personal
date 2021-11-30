package adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

const defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{ range .Paragraphs }}
    <p>{{.}}</p>
    {{ end }}
    <ul>
      {{ range .Options }}
      <li>
        <a href="/{{.Chapter}}">{{.Text}}</a>
      </li>
      {{ end }}
    </ul>
  </body>
</html>`

type handler struct{ s Story }

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))

	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func JSONStory(r io.Reader) (Story, error) {
	var story Story
	if err := json.NewDecoder(r).Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
