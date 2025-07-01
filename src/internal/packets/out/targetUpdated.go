package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x3f
* @size: 6 bytes
* @description: Used to send the target to client.
 */
type TargetUpdatedPacket struct {
	Header           byte
	VirtualId        uint32
	HealthPercentage byte
}

func (p *TargetUpdatedPacket) getHeader() byte {
	return 0x3f
}

func (p *TargetUpdatedPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}
	p.VirtualId, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.HealthPercentage, err = reader.ReadByte()
	if err != nil {
		return err
	}
	return nil
}
