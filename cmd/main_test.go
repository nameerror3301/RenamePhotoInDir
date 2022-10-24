package main

import (
	"testing"
)

func TestRenameRandom(t *testing.T) {
	t.Run("check-RenameRandom-valid", func(t *testing.T) {
		path := "/home/onebyteforlife/go/src/RenamePhotoInDir/Wallpapers"
		if err := RenameRandom(path); err != nil {
			t.Error("RenameRandom() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-RenameRandom-invalid", func(t *testing.T) {
		path := ""
		if err := RenameRandom(path); err == nil {
			t.Error("RenameRandom() - Didn't pass the test with the wrong data")
		}
	})
}

// Test func isImage()
func TestIsImage(t *testing.T) {
	t.Run("check-IsImage-valid", func(t *testing.T) {
		var i int = 0
		var shareds = []string{"one.jpg", "two.jpeg", "three.png"}
		for _, val := range shareds {
			status := isImage(val)
			if status {
				i++
			}
		}
		if i != 3 {
			t.Error("isImege() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-IsImage-invalid", func(t *testing.T) {
		var i int = 0
		var shareds = []string{"one.jpag", "two.j1peg", "three.pn6g"}
		for _, val := range shareds {
			status := isImage(val)
			if status {
				i++
			}
		}
		if i != 0 {
			t.Error("isImege() - Didn't pass the test with the wrong data ")
		}
	})
}

func TestRenameNormal(t *testing.T) {
	t.Run("check-ReadFileDir-valid", func(t *testing.T) {
		path := "/home/onebyteforlife/go/src/RenamePhotoInDir/Wallpapers"
		if err := RenameNormal(path); err != nil {
			t.Error("RenameNormal() - Did not pass the test with the correct data")
		}
	})

	t.Run("check-ReadFileDir-invalid", func(t *testing.T) {
		path := ""
		if err := RenameNormal(path); err == nil {
			t.Error("RenameNormal() - Didn't pass the test with the wrong data")
		}
	})
}
