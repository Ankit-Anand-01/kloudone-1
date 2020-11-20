package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Beginning() {
	// Read Write Mode
	file, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteAt([]byte{'G'}, 0) // Write at 0 beginning
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}

func ReadFile() {
	data, err := ioutil.ReadFile("./test1.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Content: %s", data)
}

func main() {
	fmt.Printf("\n######## Read file #########\n")
	ReadFile()

	fmt.Printf("\n\n######## WriteAt #########\n")
	Beginning()

	fmt.Printf("\n\n######## Read file #########\n")
	ReadFile()
}
