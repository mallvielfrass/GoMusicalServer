package main

import (
	"fmt"
	"log"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "search.html")

}

func main() {
	http.HandleFunc("/api", api)
	http.HandleFunc("/search", search)
	http.HandleFunc("/", search)
	fs := http.FileServer(http.Dir("opus"))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/opus/", http.StripPrefix("/opus/", fs))
	fmt.Println("Server is listening...", "\n", "localhost:5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
