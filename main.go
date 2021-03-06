package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mallvielfrass/GoMusicalServer/mod/jql"

	mod "github.com/mallvielfrass/GoMusicalServer/mod"
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
type Musica struct {
	ID string `json:"id"`
	//OwnedID string `json:"owner_id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	URL    string `json:"url"`
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "enctype")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Host)
	fmt.Println(r.URL.Path)
	key := r.URL.Query()
	fmt.Println(key)
	value, ok := key["q"]
	if ok {
		fmt.Println("value: ", value)
		mname := string(key["q"][0]) //get music name

		//res := "audio nome: " + mname
		//________________________________
		ready, err := ioutil.ReadFile("key.txt") // Считываем ключ из файла
		if err != nil {
			log.Fatalln("Не удалось считать ключ из файла:", err)
		}
		keys := (strings.Split(string(ready), "\n"))[0]
		api := pretiumvkgo.OldAPI(keys)

		fmt.Println("enter music name")
		resVK := api.AudioSearch(mname, 100, 0)

		//
		//fmt.Println(resVK)
		bx := []byte(resVK)
		var result MResult

		err = json.Unmarshal(bx, &result)
		check(err)
		R := result.Response.Items
		fmt.Println("music", R[0])
		lenR := len(R)

		//fullMusic := "title:" + R[0].Title + "<br>" + "Artist:" + R[0].Artist + "<br>" + "url: " + R[0].URL
		var Musa = []Musica{}
		i := 0
		for i < lenR {
			var msg = new(Musica)
			msg.ID = strconv.Itoa(R[i].ID)
			//msg.OwnedID = strconv.Itoa(R[i].OwnedID)
			msg.Title = R[i].Title
			msg.Artist = R[i].Artist
			msg.URL = R[i].URL
			Musa = append(Musa, *msg)
			i = i + 1
		}
		//fmt.Println(fullMusic)
		fmt.Println("Musa")
		jsMusa, err := json.Marshal(Musa)
		check(err)

		fmt.Fprintf(w, string(jsMusa))
	} else {
		fmt.Println("key q not found")
	}
	value, ok = key["link"]
	if ok {

		arr := strings.Split(value[0], "cut=")
		fmt.Println(arr)
		id, err := strconv.Atoi(arr[1])
		check(err)
		fmt.Println("id ", id)
		fmt.Println("name ", arr[2])
		fmt.Println("link ", arr[3])
		addr := jql.Search(id)
		//fmt.Println("localsscan: %s", addr)
		name := "_split_" + string(id) + "_split_" + arr[2]
		link := arr[3]

		//fmt.Println("func Download")
		fmt.Println(name, "\n", link)
		//nameDown := "music/" + name + ".mp3"
		nameOpus := "opus/" + name + ".opus"
		//if addr == "true" {
		//	fmt.Println("song is true")
		//	}
		if addr == "false" {
			fmt.Println("func Download start")
			convert(link, nameOpus)
			hash := mod.GetHash(nameOpus)
			name := mod.Rename(nameOpus, hash)
			fmt.Println("id %s name %s ", id, arr[2])
			jql.Add(id, hash, arr[2])
			http.Redirect(w, r, "/opus/"+name, 301)
		} else {
			http.Redirect(w, r, "/"+addr+".opus", 301)
		}

	} else {
		fmt.Println("key link not found")
	}
}
