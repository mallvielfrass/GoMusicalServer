package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	pretiumvkgo "github.com/mallvielfrass/pretiumVKgo"
)

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

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
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Url    string `json:"url"`
}

func api(w http.ResponseWriter, r *http.Request) {

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
		api := pretiumvkgo.NewAPI(keys)
		fmt.Println("enter music name")

		resVK := api.AudioSearch(mname, 10, 0)
		//fmt.Println(x)
		//
		//fmt.Println(resVK)
		bx := []byte(resVK)
		var result MResult

		err = json.Unmarshal(bx, &result)
		if err != nil {
			log.Fatalln(err)
		}
		R := result.Response.Items

		lenR := len(R)

		//fullMusic := "title:" + R[0].Title + "<br>" + "Artist:" + R[0].Artist + "<br>" + "url: " + R[0].URL
		var Musa = []Musica{}
		i := 0
		for i < lenR {
			var msg = new(Musica)
			msg.Title = R[i].Title
			msg.Artist = R[i].Artist
			msg.Url = R[i].URL
			Musa = append(Musa, *msg)
			i = i + 1
		}
		//fmt.Println(fullMusic)
		//fmt.Println(Musa)
		jsMusa, err := json.Marshal(Musa)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		//fmt.Println(string(jsMusa)) //musica
		//filename := "audio/" + R[0].Title + ".mp3"
		fmt.Fprintf(w, string(jsMusa))
	} else {
		fmt.Println("key q not found")
	}
	value, ok = key["link"]
	if ok {

		//fmt.Println(value[0])
		arr := strings.Split(value[0], "cut=")
		name := arr[1]
		link := arr[2]
		fmt.Println(name, "\n", link)
		nameDown := "music/" + name + ".mp3"
		nameOpus := "opus/" + name + ".opus"
		downloadFile(nameDown, link)
		command := "ffmpeg -i " + "'" + nameDown + "'" + " -c:a libopus -b:a 48k -vbr on -compression_level 10 -frame_duration 60 -application voip " + "'" + nameOpus + "'"
		fmt.Println(command)
		out, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		fmt.Println(string(out))
		http.Redirect(w, r, "/"+nameOpus, 301)

	} else {
		fmt.Println("key not found")
	}
}
func search(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "search.html")

}
func main() {
	http.HandleFunc("/api", api)
	http.HandleFunc("/search", search)
	fs := http.FileServer(http.Dir("opus"))
	http.Handle("/opus/", http.StripPrefix("/opus/", fs))
	fmt.Println("Server is listening...", "\n", "localhost:8181")
	http.ListenAndServe("localhost:8181", nil)

}
