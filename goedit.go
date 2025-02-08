package main

import (
	"fmt"
	"log"
	"os"
)

func displayFileContents(fileName string) string{
	file, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}

func newValues(){

}

func main() {
	fmt.Println("Welcome to dedit")
	displayFileContents("test.txt")
}
