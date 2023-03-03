package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Title    string `json:"title_english"`
	Synopsis string `json:"synopsis"`
	Favorite bool   `json:"favorite"`
}

type Anime struct {
	Data []Data `json:"data"`
}

func main() {
	http.HandleFunc("/", listAnimeHandler)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func listAnimeHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve data from 3rd party API
	resp, err := http.Get("https://api.jikan.moe/v4/top/anime?limit=10&filter=bypopularity")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Parse response body to Anime struct array
	var animeList Anime
	err = json.NewDecoder(resp.Body).Decode(&animeList)
	if err != nil {
		http.Error(w, "Failed to parse data", http.StatusInternalServerError)
		return
	}

	// Render HTML template
	tmpl := template.Must(template.ParseFiles("template/list.html"))
	err = tmpl.Execute(w, animeList.Data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
