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
	c := getLinesChannel(f)
	for i := range c {
		fmt.Printf("read %s\n", i)

	}
	// var noOfLines = 0
	// defer f.Close()

}

func getLinesChannel(f io.ReadCloser) chan string {
	c := make(chan string)
	var line string
	bytes8 := make([]byte, 8, 8)
	currentLine := 0
	var parts []string
	go func() {
		defer f.Close()
		defer close(c)
		for {
			n, err := f.Read(bytes8)
			if err != nil {
				if errors.Is(err, io.EOF) {
					// c <- parts[currentLine]
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
				// fmt.Printf("read :%s\n", parts[currentLine])
				c <- parts[currentLine]
				currentLine += 1
			}
		}
		c <- parts[currentLine]

	}()
	return c
}
