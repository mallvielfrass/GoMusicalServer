package main

import (
	"fmt"
	"log"
	"os/exec"
)

func convert(link, nameOpus string) {

	command := "ffmpeg -i " + "'" + link + "'" + " -c:a libopus -b:a 48k -vbr on -compression_level 10 -frame_duration 60 -application voip " + "'" + nameOpus + "'"
	fmt.Println(command)
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
		panic("some error found")
	}
	fmt.Println(string(out))
}
