package in

import (
	"encoding/binary"
	"go-mt2/internal/packets"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type HandshakePacket struct {
	Header      ClientPacketType
	DwHandshake uint32 // DWORD -> uint32 (32-bit unsigned)
	DwTime      uint32 // DWORD -> uint32 (32-bit unsigned)
	LDelta      int32  // LONG -> int32 (32-bit signed)
}

func (p *HandshakePacket) GetHeader() byte {
	return byte(p.Header)
}

func (p *HandshakePacket) GetType() byte {
	return 255
}

func (p *HandshakePacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.DwHandshake, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.DwTime, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.LDelta, err = reader.ReadInt32()
	if err != nil {
		return err
	}
	return nil
}

func HandleHandshake(packet *HandshakePacket, conn *websocket.Conn) error {
	log.Println("Handshake packet received")

	data := make([]byte, 1+4+4+4)

	data[0] = byte(HEADER_CG_HANDSHAKE)
	dwHandshake := packet.DwHandshake
	dwTime := uint32(time.Now().Unix())
	lDelta := int32(packet.DwTime) - int32(dwTime)

	binary.LittleEndian.PutUint32(data[1:5], uint32(dwHandshake))
	binary.LittleEndian.PutUint32(data[5:9], uint32(dwTime))
	binary.LittleEndian.PutUint32(data[9:13], uint32(lDelta))

	conn.WriteMessage(websocket.BinaryMessage, data)
	return nil

}
