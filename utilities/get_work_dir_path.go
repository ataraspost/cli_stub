package utilities

import (
	"fmt"
	"os"
	"strings"
)

func GetWorkDirPath(path string) string {

	if path == "." {
		work_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		return work_dir
	} else if strings.HasPrefix(path, ".") {
		work_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		work_dir = work_dir + path[1:]
		return work_dir
	} else if strings.HasPrefix(path, "/") {
		work_dir := path
		return work_dir
	} else {
		work_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		work_dir = work_dir + "/" + path
		return work_dir
	}
}
