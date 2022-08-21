package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type Note struct {
	Data     []byte
	Destruct bool
}

type Server struct {
	BaseURL string
	cache   *cache.Cache
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET" || r.Method == "HEAD":
		if path := r.URL.Path; path != "/" {
			note, noteID := &Note{}, strings.TrimPrefix(path, "/")
			if err := s.cache.
				GetSkippingLocalCache(r.Context(), noteID, note); err != nil {
				s.badRequest(
					w,
					http.StatusNotFound,
					fmt.Sprintf("note with ID %q does not exist.", noteID),
				)
				return
			}
			if note.Destruct {
				if err := s.cache.Delete(r.Context(), noteID); err != nil {
					fmt.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Something went wrong. Please check the server logs."))
				}
			}
			w.Write(note.Data)
			return
		}
		s.renderTemplate(w, r, nil, "layout", "dist/layout.html", "dist/index.html")
		return
	case r.Method == "POST" || r.URL.Path == "/":
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			s.badRequest(
				w, http.StatusUnsupportedMediaType, "Invalid media type posted.")
			return
		}
		if err := r.ParseForm(); err != nil {
			s.badRequest(w, http.StatusBadRequest, "Invalid form data posted.")
		}
		form := r.PostForm
		destruct := false
		ttl := time.Hour * 24
		if form.Get("ttl") == "untilRead" {
			destruct = true
			ttl = ttl * 365
		}

		key := uuid.NewString()
		if err := s.cache.Set(
			&cache.Item{
				Ctx:            r.Context(),
				Key:            key,
				Value:          &Note{Data: []byte(form.Get("message")), Destruct: destruct},
				TTL:            ttl,
				SkipLocalCache: true,
			},
		); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong. Please check the server logs."))
			return
		}

		noteURL := fmt.Sprintf("%s/%s", s.BaseURL, key)
		s.renderTemplate(
			w, r,
			struct {
				Title      string
				Paragraphs []any
			}{
				Title:      "Note was successfully created",
				Paragraphs: []any{template.HTML(fmt.Sprintf("<a href='%s'>%s</a>", noteURL, noteURL))},
			},
			"layout", "dist/layout.html", "dist/message.html",
		)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

func (s Server) badRequest(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}

func (s *Server) renderTemplate(
	w http.ResponseWriter,
	r *http.Request,
	data any,
	name string,
	files ...string,
) {
	if err := template.Must(template.ParseFiles(files...)).
		ExecuteTemplate(w, name, data); err != nil {
		log.Fatal(err)
	}
}

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if len(redisURL) == 0 {
		redisURL = "redis://:@localhost:6379/1"
	}
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(opts)
	defer client.Close()
	log.Fatal(http.ListenAndServe(":9001", Server{
		cache: cache.New(&cache.Options{
			Redis: client,
		}),
	}))
}
