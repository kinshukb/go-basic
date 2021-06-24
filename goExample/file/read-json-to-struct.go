package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {

	type FileAttr struct {
		Id       int
		Username string
		Hostname string
		Cpu      float64
		Mem      float64
		Version  string
	}

	output := FileAttr{}

	userFile, err := os.Open("user.json")
	if err != nil {
		log.Println("opening config file", err.Error())
	}
	jsonParser := json.NewDecoder(userFile)
	if err = jsonParser.Decode(&output); err != nil {
		log.Println("parsing config file", err.Error())
	}
	log.Println("Contents of file as go struct: ", output)
}
