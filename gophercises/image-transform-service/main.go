package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"prim/primitive"

	"golang.org/x/image/webp"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><body>
	<form action="upload" method="post" enctype="multipart/form-data">
		<input type="file" name="image" />
		<input type="submit" text>Upload image</input>
	</form>
</body></html>`))
	})
	mux.HandleFunc(http.MethodPost+" /upload", func(w http.ResponseWriter, r *http.Request) {
		f, hdr, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()

		raw, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		modes := []primitive.Mode{
			primitive.ModeCombo,
			primitive.ModeTriangle,
			primitive.ModeRotatedEllipse,
			primitive.ModeCircle,
			primitive.ModePolygon,
			primitive.ModeRotatedRect,
		}
		images := make([]string, len(modes))
		start := time.Now()
		var wg sync.WaitGroup
		for i, mode := range modes {
			i := i
			wg.Add(1)
			go func(mode primitive.Mode) {
				defer wg.Done()
				var buf bytes.Buffer
				ext := filepath.Ext(hdr.Filename)
				if ext == ".webp" {
					i, err := webp.Decode(bytes.NewReader(raw))
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					if err := jpeg.Encode(&buf, i, nil); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					ext = ".jpeg"
				} else {
					if _, err := buf.Write(raw); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
				out, err := primitive.Transform(
					r.Context(),
					primitive.ImageFD{
						Reader: &buf,
						Ext:    ext,
					},
					220,
					mode,
				)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fname := outFile(hdr.Filename, mode)
				imgF, err := os.Create("./img/" + fname)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				if _, err := io.Copy(imgF, out); err != nil {
					log.Println(err)
				}
				images[i] = fname
			}(mode)
		}
		wg.Wait()
		fmt.Println("Took ", start.Sub(time.Now()).Abs().String())

		if err := template.Must(template.New("").Parse(`<html><body>
{{ range . }}
<img src="/img/{{.}}" />
{{ end }}
</body></html>`)).Execute(w, images); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))

	if err := http.ListenAndServe("127.0.0.1:8080", mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func outFile(name string, mode primitive.Mode) (outName string) {
	ext := filepath.Ext(name)
	return "out_" + name[:len(name)-len(ext)] + "_" + mode.String() + ext
}
