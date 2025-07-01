package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x11
* @size: 22 bytes
* @description: Is used to send and update of a point (attribute) to the client (used to notify all nearby players of an update of a character). See all points in PointsEnum.
 */
type CharacterPointChangePacket struct {
	Header byte
	Vid    uint32
	Type   byte
	Amount uint64
	Value  uint64
}

func (p *CharacterPointChangePacket) getHeader() byte {
	return 0x11
}

func (p *CharacterPointChangePacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}

	p.Vid, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.Type, err = reader.ReadByte()
	if err != nil {
		return err
	}

	p.Amount, err = reader.ReadUInt64()
	if err != nil {
		return err
	}

	p.Value, err = reader.ReadUInt64()
	if err != nil {
		return err
	}
	return nil
}
