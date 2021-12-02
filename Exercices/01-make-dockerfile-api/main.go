package main

import (
	"fmt"
	"net/http"

	b64 "encoding/base64"

	"github.com/go-chi/chi"
)

func resultadoEjercicio(w http.ResponseWriter, r *http.Request) {
	cData := "PGgxPkZlbGljaXRhY2lvbmVzIGNvbXBsZXRhc3RlIGVsIGVqZXJjaWNpbyEhITwvaDE-"
	uDec, _ := b64.URLEncoding.DecodeString(cData)
	w.Write([]byte(string(uDec)))
}

func main() {
	fmt.Println("API Alumnos in GO")
	fmt.Println("Running on http://0.0.0.0:5000/")
	r := chi.NewRouter()
	r.Get("/", resultadoEjercicio)

	http.ListenAndServe(":5000", r)
}
