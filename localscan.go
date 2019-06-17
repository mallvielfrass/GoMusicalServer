package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func localscan(id string) string {

	files, err := ioutil.ReadDir("opus")
	check(err)

	ret := "false"
	for _, file := range files {
		name := file.Name()
		idz := strings.Split(name, "_split_")
		//fmt.Println(id, " ", idz[1])
		if idz[1] == id {
			fmt.Println("song is true", name)
			ret = "true"
		}
	}
	return ret
}
