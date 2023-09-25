package common

import (
	"os"
	"log"
)

func IsFileOrDir(path string) (string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Panicf("Error while opening %s: %s", path, err)
	}

	if fileInfo.Mode().IsRegular() {
		return "file"
	} else if fileInfo.IsDir() {
		return "dir"
	} else {
		log.Panicf("Error: %s is not a file nor directory: %s", path, err)
		return ""
	}	
}
