package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("exit")
			break
		}
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Send %d bytes data\n", n)
	}
}
