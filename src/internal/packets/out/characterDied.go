package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x0e
* @size: 5 bytes
* @description: Used to notify the client when some entity has died.
 */
type CharacterDiedPacket struct {
	Header    byte
	VirtualId uint32
}

func (p *CharacterDiedPacket) getHeader() byte {
	return 0x0e
}

func (p *CharacterDiedPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}

	p.VirtualId, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	return nil
}
