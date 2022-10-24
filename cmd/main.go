package main

import (
	"fmt"
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

	if err := RenameNormal(path); err != nil {
		log.Fatalf("%s", err)
	}
}

// Recursive search and renaming files to normal names
func RenameNormal(path string) error {
	var idx int
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Err path is not valid, your path <%s> - %s", path, err)
	}
	for _, file := range files {
		if file.IsDir() {
			RenameNormal(filepath.Join(path, file.Name()))
		} else {
			if isImage(file.Name()) {
				idx++
				shared := filepath.Ext(filepath.Join(path, file.Name()))
				if err := os.Rename(filepath.Join(path, file.Name()), filepath.Join(path, fmt.Sprintf("Images-%s%s", strconv.Itoa(idx), shared))); err != nil {
					return fmt.Errorf("Err rename file [%s] in dir [%s] - %s", file.Name(), path, err)
				}
			}
		}
	}
	return nil
}

// Initial renaming of files to random names
func RenameRandom(path string) error {
	if _, err := os.ReadDir(path); err != nil {
		return fmt.Errorf("Err path is not valid, your path <%s> - %s", path, err)
	}
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if !info.IsDir() && isImage(info.Name()) {
			if err != nil {
				return err
			}
			shared := filepath.Ext(wPath)
			dir, file := filepath.Split(wPath)
			if err := os.Rename(wPath, filepath.Join(dir, fmt.Sprintf("%s%s", GenRandomString(), shared))); err != nil {
				return fmt.Errorf("Err rename file [%s] in dir [%s] - %s", file, dir, err)
			}
			return nil
		}
		return nil
	})
	return nil
}

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

func isImage(file string) bool {
	var shareds = []string{".jpg", ".jpeg", ".png"}
	for _, shared := range shareds {
		if filepath.Ext(file) == shared {
			return true
		}
	}
	return false
}
