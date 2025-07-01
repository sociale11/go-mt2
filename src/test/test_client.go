package main

import (
	"go-mt2/internal/packets/in"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to your server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Create a login packet
	loginData := createLoginPacket("admin", "admin")

	for i := 0; i < 10; i++ {
		log.Printf("Sending: %x", loginData)

		// Send the packet
		err = conn.WriteMessage(websocket.BinaryMessage, loginData)
		if err != nil {
			log.Println("write:", err)
			return
		}
	}

	// Read response
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	log.Printf("Received: %x", message)
}

func createLoginPacket(username, password string) []byte {
	data := make([]byte, 1+31+17) // header + username + password

	// Header
	data[0] = byte(in.HEADER_CG_LOGIN)

	// Username (31 bytes, null-padded)
	copy(data[1:32], username)

	// Password (17 bytes, null-padded)
	copy(data[32:49], password)

	return data
}
