package main

import (
	"encoding/binary"
	"go-mt2/internal/packets/in"
	"log"
	"time"

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
	_ = createLoginPacket("admin", "admin")

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		if data[0] == byte(in.HEADER_CG_HANDSHAKE) {
			conn.WriteMessage(websocket.BinaryMessage, createHandshakePacket(data))
		}
	}

	// for range 100 {
	// log.Printf("Sending: %x", loginData)

	// Send the packet
	// 	err = conn.WriteMessage(websocket.BinaryMessage, handshake)
	// 	if err != nil {
	// 		log.Println("write:", err)
	// 		return
	// 	}
	// }

	// Read response
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	log.Printf("Received: %x", message)
}

func createLoginPacket(username, password string) []byte {
	data := make([]byte, 1+30+16) // header + username + password
	data[0] = byte(in.HEADER_CG_LOGIN)
	copy(data[1:31], username)
	copy(data[32:47], password)
	return data
}

func createHandshakePacket(data []byte) []byte {
	newPacket := make([]byte, 1+4+4+4)

	data[0] = byte(in.HEADER_CG_HANDSHAKE)
	dwHandshake := binary.LittleEndian.Uint32(data[1:5])
	dwTime := uint32(time.Now().Unix())
	lDelta := int32(binary.LittleEndian.Uint32(data[9:13])) - int32(dwTime)

	binary.LittleEndian.PutUint32(newPacket[1:5], uint32(dwHandshake))
	binary.LittleEndian.PutUint32(newPacket[5:9], uint32(dwTime))
	binary.LittleEndian.PutUint32(newPacket[9:13], uint32(lDelta))

	return data
}
