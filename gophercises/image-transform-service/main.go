package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"prim/primitive"
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

		fmt.Println(hdr.Filename)
		out, err := primitive.Transform(r.Context(), primitive.ImageFD{
			Reader: f,
			Ext:    filepath.Ext(hdr.Filename),
		}, 40)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		if _, err := io.Copy(w, out); err != nil {
			log.Println(err)
		}
	})

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
