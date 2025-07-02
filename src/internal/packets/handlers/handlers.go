package handlers

import (
	"fmt"
	"go-mt2/internal/packets"
	"go-mt2/internal/packets/in"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:   255,
	WriteBufferSize:  255,
	HandshakeTimeout: 10 * time.Second,
}

func ReadPacket(w http.ResponseWriter, r *http.Request) {
	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for {
		_, data, err := c.ReadMessage()
		_, err = CreatePacketFromData(data)

		if err != nil {
			log.Println("read:", err)
			continue
		}

	}
}
func CreatePacketFromData(data []byte) (in.ClientPacket, error) {

	if len(data) < 1 {
		return nil, fmt.Errorf("no data to parse")
	}

	packetType := in.ClientPacketType(data[0])

	switch packetType {
	case in.HEADER_CG_LOGIN:
		{
			packet := &in.LoginPacket{}
			packet.Header = in.ClientPacketType(packet.GetType())
			br := packets.NewBufferReader(data, 1)
			err := packet.Parse(br)
			if err != nil {
				return nil, fmt.Errorf("failed to parse login packet: %w", err)
			}
			return packet, nil
		}
	case in.HEADER_CG_HANDSHAKE:
		{

		}
	default:
		return nil, fmt.Errorf("unknown packet type: %d", packetType)
	}
}
