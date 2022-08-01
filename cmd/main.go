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
	path := "your path to files photo"
	if status := RenameRandom(path); !status {
		log.Fatal("Err random rename file")
	}

	if err := RenameFileInDir(path); err != nil {
		log.Fatal(err)
	}
}

// A final function that combines all functions and checks for a directory
func RenameFileInDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("The directory does not exist - %s", err)
	}
	if err := ReadFileDir(path); err != nil {
		return err
	}
	return nil
}

// Recursive search and renaming files to normal names
func ReadFileDir(path string) error {
	var idx int
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			ReadFileDir(filepath.Join(path, file.Name()))
		} else {
			if isImage(file.Name()) {
				idx++
				shared := filepath.Ext(filepath.Join(path, file.Name()))
				os.Rename(filepath.Join(path, file.Name()), filepath.Join(path, fmt.Sprintf("Images-%s%s", strconv.Itoa(idx), shared)))
			}
		}
	}
	return nil
}

// Initial renaming of files to random names
func RenameRandom(path string) bool {
	var pathToFile string
	filepath.Walk(path, func(wPath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			pathToFile = wPath
		}
		if isImage(wPath) {
			shared := filepath.Ext(wPath)
			err := os.Rename(wPath, filepath.Join(pathToFile, fmt.Sprintf("%s%s", GenRandomString(), shared)))
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
	return true
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
