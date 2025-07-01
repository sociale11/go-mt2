package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x87
* @size: 10 bytes
* @description: Used to send the damage to client.
 */
type DamagePacket struct {
	Header      byte
	VirtualId   uint32
	DamageFlags byte
	Damage      uint32
}

func (p *DamagePacket) getHeader() byte {
	return 0x87
}

func (p *DamagePacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}
	p.VirtualId, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.DamageFlags, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Damage, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	return nil
}
