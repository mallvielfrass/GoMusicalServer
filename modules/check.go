package modules

//на будущее
import "os"

func check() string {
	if _, err := os.Stat("../opus"); os.IsNotExist(err) {
		return "false. create dir opus please,and restart programm"
	}
	return "true"
}
