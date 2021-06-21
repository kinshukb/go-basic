package main

import (
	"encoding/json"
	"io/ioutil"
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

	input := FileAttr{
		Id:       101,
		Username: "king",
		Hostname: "odcorp",
		Cpu:      4.565,
		Mem:      5.901,
		Version:  "0.4",
	}

	// convert your struct into []byte using json package
	newFile, _ := json.MarshalIndent(input, "", " ")
	// use the []byte generated above to fill it in your file
	_ = ioutil.WriteFile("user.json", newFile, 0755)
}
