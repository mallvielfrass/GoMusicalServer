package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pretiumvkgo "github.com/mallvielfrass/pretiumVKgo"
)

type MResult struct {
	Response MResponse `json:"response"`
}

type MResponse struct {
	Count int     `json:"count"`
	Items []MItem `json:"items"`
}
type MItem struct {
	ID       int    `json:"id"`
	OwnedID  int    `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Date     int    `json:"date"`
	URL      string `json:"url"`
}

func api(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)

	fmt.Println(r.URL)
	fmt.Println(r.URL.Host)
	fmt.Println(r.URL.Path)
	key := r.URL.Query()
	mname := string(key["q"][0]) //get music name
	//res := "audio nome: " + mname
	//________________________________
	ready, err := ioutil.ReadFile("key.txt") // Считываем ключ из файла
	if err != nil {
		log.Fatalln("Не удалось считать ключ из файла:", err)
	}
	keys := (strings.Split(string(ready), "\n"))[0]
	api := pretiumvkgo.NewAPI(keys)
	fmt.Println("enter music name")

	resVK := api.AudioSearch(mname, 10, 0)
	//fmt.Println(x)
	bx := []byte(resVK)
	var result MResult

	err = json.Unmarshal(bx, &result)
	if err != nil {
		log.Fatalln(err)
	}
	R := result.Response.Items
	fmt.Println(len(R))
	fullMusic := "title:" + R[0].Title + "<br>" + "Artist:" + R[0].Artist + "<br>" + "url: " + R[0].URL
	fmt.Println(fullMusic)
	//filename := "audio/" + R[0].Title + ".mp3"
	fmt.Fprintf(w, fullMusic)

}
func search(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "search.html")

}
func main() {
	http.HandleFunc("/api", api)
	http.HandleFunc("/search", search)
	fmt.Println("Server is listening...", "\n", "localhost:8181")
	http.ListenAndServe("localhost:8181", nil)

}
