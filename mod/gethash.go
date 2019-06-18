package mod

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func GetHash(patch string) string {
	f, err := os.Open(patch)
	check(err)
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("len=", len(hash))
	fmt.Println(hash)
	return hash
}
func Rename(nameOpus, hash string) string {

	newName := hash + ".opus"

	err := os.Rename(nameOpus, "opus/"+newName)
	check(err)
	return newName
}
