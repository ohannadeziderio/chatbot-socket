package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("You: ")
		message, _ := reader.ReadString('\n')

		conn.Write([]byte(message))

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Cannot read the data:", err)
			break
		}

		fmt.Println("ChatBot:", string(buf[:n]))
	}

	conn.Close()
}
