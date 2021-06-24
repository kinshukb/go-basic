package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	quickRead()
	bufferedReader()
}

func quickRead() {
	// Read file to byte slice
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}

func bufferedReader() {
	// Open file and create a buffered reader on top
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Data as hex: %x\n", data)
	// fmt.Printf("Data as string: %s\n", data)
	// fmt.Println("Number of bytes read:", len(data))

	bufferedReader := bufio.NewReader(file)

	// Get bytes without advancing pointer
	byteSlice := make([]byte, 10)
	byteSlice, err = bufferedReader.Peek(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// Read and advance pointer
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// Ready 1 byte. Error if no byte to read
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// Read up to and including delimiter
	// Returns byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

	// Read up to and including delimiter
	// Returns string
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)

	// This example reads a few lines so test.txt
	// should have a few lines of text to work correct
}
