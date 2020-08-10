package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Alumno struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var alumnos = make([]Alumno, 0, 0)

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	log.Printf("[JSON Response] %v", payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllAlumnos(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, http.StatusOK, alumnos)
}

func AddAlumnos(w http.ResponseWriter, r *http.Request) {
	var alumno Alumno
	json.NewDecoder(r.Body).Decode(&alumno)
	idTemp := len(alumnos) + 1
	alumno.Id = idTemp
	alumnos = append(alumnos, alumno)
	responseJSON(w, http.StatusOK, alumno)
}

func main() {
	fmt.Println("API Alumnos in GO")
	fmt.Println("Running on http://0.0.0.0:5000/")
	r := chi.NewRouter()
	r.Get("/alumnos", AllAlumnos)
	r.Post("/alumnos", AddAlumnos)

	http.ListenAndServe(":5000", r)
}
