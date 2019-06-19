package jql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type VkMusic struct {
	ID   int
	VkID int
	Hash string
	Name string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Search(VarVkID int) string {
	db, err := sql.Open("sqlite3", "./VkMusic.db")
	check(err)
	defer db.Close()
	fmt.Println("VarVkID", VarVkID)
	rows, err := db.Query("SELECT * FROM Music WHERE VkID=$1", VarVkID)
	if err != nil {
		fmt.Println("Error!")
	}

	//fmt.Println(numbe)
	musDate := []VkMusic{}
	//fmt.Println(rows.Next())
	for rows.Next() {
		p := VkMusic{}
		err := rows.Scan(&p.ID, &p.VkID, &p.Hash, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		musDate = append(musDate, p)
	}
	lemD := len(musDate)
	ret := "false"
	fmt.Println("len:", lemD)
	if lemD != 0 {
		ret = musDate[0].Hash
		fmt.Println(musDate[0].ID, musDate[0].VkID, musDate[0].Hash, musDate[0].Name)
	}
	
	return ret
}
func Add(VarVkID int, VarHash, VarName string) {

	db, err := sql.Open("sqlite3", "./VkMusic.db")
	check(err)
	defer db.Close()

	rows, err := db.Query("SELECT count(*)  FROM Music")
	if err != nil {
		fmt.Println("Error!")
	}
	var numbe int64
	numbe = 0
	//fmt.Println(numbe)
	for rows.Next() {
		var count int64
		rows.Scan(&count)
		numbe = count + 1
		//fmt.Println("count", numbe)
	}

	_, err = db.Exec("insert into Music (ID, VkID, Hash, Name) values ($1, $2, $3, $4)",
		numbe, VarVkID, VarHash, VarName)
	check(err)

	//	fmt.Println(count) // id последнего добавленного объекта
	//fmt.Println(result.RowsAffected()) // количество добавленных строк

}
func Read() {
	db, err := sql.Open("sqlite3", "./VkMusic.db")
	check(err)
	defer db.Close()
	rows, err := db.Query("select * from Music")
	check(err)
	defer rows.Close()
	musDate := []VkMusic{}

	for rows.Next() {
		p := VkMusic{}
		err := rows.Scan(&p.ID, &p.VkID, &p.Hash, &p.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		musDate = append(musDate, p)
	}
	fmt.Println("len:", len(musDate))
	for _, p := range musDate {
		fmt.Println(p.ID, p.VkID, p.Hash, p.Name)
	} //fmt.Println(result.LastInsertId())

}
func Create() {
	db, err := sql.Open("sqlite3", "./VkMusic.db")
	check(err)
	defer db.Close()

	sqlStmt := `
	CREATE TABLE Music(
		ID		INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
		VkID	INTEGER,
		Hash	TEXT,
		Name	TEXT
	  )
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}
func Mainez() {
	//os.Remove("./foo.db")
	if _, err := os.Stat("./VkMusic.db"); err == nil {
		Add(129346, "AnyHash", "Song")

	} else if os.IsNotExist(err) {
		Create()

	}
	Read()
}
