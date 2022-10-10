package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	path := "/home/onebyteforlife/go/src/RenamePhotoInDir/Wallpapers"
	if err := RenameRandom(path); err != nil {
		log.Fatalf("%s", err)
	}

	if err := RenameFileInDir(path); err != nil {
		log.Fatalf("%s", err)
	}

}

// A final function that combines all functions and checks for a directory
func RenameFileInDir(path string) error {
	if err := ReadFileDir(path); err != nil {
		return fmt.Errorf("err rename file in dir - %s", err)
	}
	return nil
}

// Recursive search and renaming files to normal names
func ReadFileDir(path string) error {
	var idx int
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("err read dir - %s", err)
	}
	for _, file := range files {
		if file.IsDir() {
			ReadFileDir(filepath.Join(path, file.Name()))
		} else {
			if isImage(file.Name()) {
				idx++
				shared := filepath.Ext(filepath.Join(path, file.Name()))
				if err := os.Rename(filepath.Join(path, file.Name()), filepath.Join(path, fmt.Sprintf("Images-%s%s", strconv.Itoa(idx), shared))); err != nil {
					return fmt.Errorf("err rename file - %s", err)
				}
			}
		}
	}
	return nil
}

// Initial renaming of files to random names
func RenameRandom(path string) error {
	if _, err := os.ReadDir(path); err != nil {
		return err
	}
	var pathToFile string
	filepath.Walk(path, func(wPath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			pathToFile = wPath
		}
		if isImage(wPath) {
			shared := filepath.Ext(wPath)
			if err := os.Rename(wPath, filepath.Join(pathToFile, fmt.Sprintf("%s%s", GenRandomString(), shared))); err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

// Generating a random string for file names
func GenRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// Function for checking if a file is an image
func isImage(file string) bool {
	var shareds = []string{".jpg", ".jpeg", ".png"}
	for _, shared := range shareds {
		if filepath.Ext(file) == shared {
			return true
		}
	}
	return false
}
