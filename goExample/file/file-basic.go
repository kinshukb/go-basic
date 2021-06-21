package main

import (
	"log"
	"os"
	"time"
)

var (
	newFile *os.File
	err     error
)

func main() {

	createEmptyFile()
	truncateFile()
	getFileInfo()
	changeFileInfo()
}

func createEmptyFile() {
	// Create Empty File
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func truncateFile() {
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}

func getFileInfo() {
	var (
		fileInfo os.FileInfo
		err      error
	)

	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File name:", fileInfo.Name())
	log.Println("Size in bytes:", fileInfo.Size())
	log.Println("Permissions:", fileInfo.Mode())
	log.Println("Last modified:", fileInfo.ModTime())
	log.Println("Is Directory: ", fileInfo.IsDir())
	log.Printf("System interface type: %T\n", fileInfo.Sys())
	log.Printf("System info: %+v\n\n", fileInfo.Sys())
}

func changeFileInfo() {
	// Change perrmissions using Linux style
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change ownership
	err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change timestamps
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
