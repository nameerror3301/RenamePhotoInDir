package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path := "path"
	err := RenameFileInDir(path)
	if err != nil {
		log.Fatal(err)
	}
}

func RenameFileInDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("The directory does not exist - %s", err)
	}
	if err := ReadFileDir(path); err != nil {
		return err
	}
	return nil
}

func ReadFileDir(path string) error {
	var idx int
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			ReadFileDir(filepath.Join(path, file.Name()))
		} else {
			if isImage(file.Name()) {
				if strings.Contains(file.Name(), "[$]") {
					idx++
					continue
				} else {
					idx++
					shared := filepath.Ext(filepath.Join(path, file.Name()))
					os.Rename(filepath.Join(path, file.Name()), filepath.Join(path, fmt.Sprintf("Ima[$]ge-%s%s", strconv.Itoa(idx), shared)))
				}
			}
		}
	}
	return nil
}

func isImage(file string) bool {
	var shareds = []string{".jpg", ".jpeg", ".png"}
	for _, shared := range shareds {
		if filepath.Ext(file) == shared {
			return true
		}
	}
	return false
}
