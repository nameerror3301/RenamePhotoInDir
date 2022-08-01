package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	path := "/home/onebyteforlife/go/src/RenamePhotoInDir/Wallpapers"
	err := RenameFileInDir(path)
	if err != nil {
		fmt.Printf("Выбраная вами директория не существует - %v\n", err)
	}
}

func RenameFileInDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	if err := ReadFileDir(path); err != nil {
		return err
	}
	return nil
}

func ReadFileDir(path string) error {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
			ReadFileDir(filepath.Join(path, file.Name()))
		} else {
			if isImage(file.Name()) {
				fmt.Printf("\t%s\n", file.Name())
			}
		}
	}
	return nil
}

func isImage(file string) bool {
	if filepath.Ext(file) == ".jpg" || filepath.Ext(file) == ".png" {
		return true
	}
	return false
}
