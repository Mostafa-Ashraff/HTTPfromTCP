package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	// fmt.Println("I hope i get the job")
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		os.Exit(1)
	}
	defer listener.Close() // Ensure listener is closed on program exit

	fmt.Println("TCP listener is running on :42069")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) //
	}
	c := getLinesChannel()
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
