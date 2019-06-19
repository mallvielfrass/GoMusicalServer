package main

import (
	"fmt"

	jql "github.com/mallvielfrass/GoMusicalServer/mod/jql"
)

func main() {
	//os.Remove("./foo.db")
	jql.Add(129346, "AnyHash", "Song")
	//	jql.Read()
	fmt.Println(jql.Search(129336))

}
