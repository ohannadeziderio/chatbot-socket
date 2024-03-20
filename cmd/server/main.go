package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ohannadeziderio/chatbot-socket/pkg/chat"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error when init tcp connection", err)
		return
	}

	bot := chat.NewChatbot()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error when accept connection", err)
			return
		}

		log.Println("Getting new connection:", conn.RemoteAddr())

		go handleConnection(conn, *bot)
	}
}

func handleConnection(conn net.Conn, bot chat.Chatbot) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Cannot read the data:", err)
			return
		}

		message := string(buf[:n])
		log.Print("Receiving message: ", conn.RemoteAddr(), " ", message)

		if message == "quit\n" {
			log.Println("Client", conn.RemoteAddr(), " closed the connection.")
			return
		}

		response := bot.Answer(message)
		log.Print("Responding to client ", conn.RemoteAddr(), " ", response)

		conn.Write([]byte(response + "\n"))
	}
}
