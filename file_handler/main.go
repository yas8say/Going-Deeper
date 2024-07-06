package main

import (
	"fmt"
	"os"
)

func main() {
	handleFile("test.txt")
	handleFile("nonexistent.txt")
}

func handleFile(filename string) {
	defer func() {
		if r := recover(); r != nil {
			//nil will be printed when there's no error
			fmt.Println("Recovered from error:", r)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %s", err))
	}

	defer file.Close()

	fmt.Println("File opened successfully", filename)
}
