package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	// fmt.Println("I hope i get the job")
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s \n", inputFilePath, err)
	}
	defer f.Close()
	bytes8 := make([]byte, 8, 8)
	var line string
	// var noOfLines = 0
	currentLine := 0
	var parts []string
	for {
		n, err := f.Read(bytes8)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s \n", err.Error())
			break
		}
		str := string(bytes8[:n])
		line += str
		parts = strings.Split(line, "\n")
		if len(parts)-1 > currentLine {
			// fmt.Printf("read :%v\n", parts)
			fmt.Printf("read :%s\n", parts[currentLine])
			currentLine += 1
		}
	}
	fmt.Printf("read :%s\n", parts[currentLine])

}
