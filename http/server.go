package main

import (
	"curcli/model"
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed build
var embeddedFiles embed.FS

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get date from query params
	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get currname from params
	curr := r.URL.Query().Get("curr")

	// find rates in database
	rates, err := model.FindRates(date, curr)
	if err == model.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(rates)
}

func getFileSystem() http.FileSystem {

	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		log.Panic(err)
	}

	return http.FS(fsys)
}

func main() {
	log.Println("Setting up database")
	err := model.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server on port 8080")
	http.Handle("/", http.FileServer(getFileSystem()))
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/api/rates", getRates)
	http.ListenAndServe(":8080", nil)
}
