package main

import (
	"bytes"
	"errors"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

		var buf bytes.Buffer
		ext := filepath.Ext(hdr.Filename)
		if ext == ".webp" {
			i, err := webp.Decode(f)
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
			if _, err := buf.ReadFrom(f); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		out, err := primitive.Transform(r.Context(), primitive.ImageFD{
			Reader: &buf,
			Ext:    ext,
		}, 200)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		imgF, err := os.Create("./img/out_" + hdr.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(imgF, out); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/img/out_"+hdr.Filename, http.StatusFound)
	})

	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))

	if err := http.ListenAndServe("127.0.0.1:8080", mux); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	// f, err := os.Open("./input.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// out, err := primitive.Transform(context.TODO(), f, 10)
	// if err != nil {
	// 	panic(err)
	// }
	// outF, err := os.Create("output.jpg")
	// if _, err = io.Copy(outF, out); err != nil {
	// 	panic(err)
	// }
}
