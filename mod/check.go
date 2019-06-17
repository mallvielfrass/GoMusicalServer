package mod

//на будущее
import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func CheckFile(name string) string {
	//check(os.Chdir("../"))

	if _, err := os.Stat(name); os.IsNotExist(err) {
		return "false. create dir opus please,and restart programm"
	}
	return "checkfile true"
}
