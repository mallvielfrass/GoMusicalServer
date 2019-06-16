package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check() string {
	if _, err := os.Stat("opus"); os.IsNotExist(err) {
		return "false. create dir opus please,and restart programm"
	}
	return "true"
}
func localscan(id string) string {

	files, err := ioutil.ReadDir("opus")
	if err != nil {
		log.Fatal(err)

	}
	ret := "false"
	for _, file := range files {
		name := file.Name()
		idz := strings.Split(name, "_split_")
		fmt.Println(id, " ", idz[1])
		if idz[1] == id {

			fmt.Println("song is true", name)
			ret = "true"
		}
	}
	return ret
}
