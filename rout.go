package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mod "github.com/mallvielfrass/GoMusicalServer/mod"
	"github.com/mallvielfrass/GoMusicalServer/mod/jql"
)

//отдельная функция для проверки ошибок. юзать:
//	exampl, err := ("lol")
//		check(err)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func search(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/search.html")
}
func apiHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/api.html")
}
func aboutHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/about.html")
}

func main() {
	//test
	if _, err := os.Stat("./VkMusic.db"); err == nil {
		//jql.Add(129346, "AnyHash", "Song")
		fmt.Println("table Ok")

	} else if os.IsNotExist(err) {
		jql.Create()

	}
	//pages
	http.HandleFunc("/", search)
	http.HandleFunc("/api_help", apiHTML)
	http.HandleFunc("/about", aboutHTML)
	http.HandleFunc("/search", search)

	//api
	http.HandleFunc("/api", api)

	//	static files
	fs := http.FileServer(http.Dir("opus"))
	static := http.FileServer(http.Dir("music/static"))

	http.Handle("/opus/", http.StripPrefix("/opus/", fs))
	http.Handle("/js/", http.StripPrefix("/js/", static))
	http.Handle("/css/", http.StripPrefix("/css/", static))

	//start
	ip := ":5050"
	fmt.Println("check folder: ", mod.CheckFile("opus"))
	fmt.Println("Server is listening...", "\n", "localhost%s", ip)
	log.Fatal(http.ListenAndServe(ip, nil))
}
