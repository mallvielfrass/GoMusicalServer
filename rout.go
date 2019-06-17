package main

import (
	"fmt"
	"log"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "search.html")
}
func api_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api.html")
}
func about_html(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about.html")
}

func main() {
	http.HandleFunc("/api_help", api_html)
	http.HandleFunc("/about", about_html)
	http.HandleFunc("/api", api)
	http.HandleFunc("/search", search)
	http.HandleFunc("/", search)
	fs := http.FileServer(http.Dir("opus"))
	static := http.FileServer(http.Dir("music/static"))
	http.Handle("/opus/", http.StripPrefix("/opus/", fs))
	http.Handle("/js/", http.StripPrefix("/js/", static))
	http.Handle("/css/", http.StripPrefix("/css/", static))
	fmt.Println("Server is listening...", "\n", "localhost:5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
