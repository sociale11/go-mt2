package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x10
* @size: 1021 bytes
* @description: Is used to send update of all the points (attributes) of a character to the client. See all points in PointsEnum.
 */
type CharacterPointsPacket struct {
	Header byte
	Points uint32
}

func (p *CharacterPointsPacket) getHeader() byte {
	return 0x10
}

func (p *CharacterPointsPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}
	p.Points, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	return nil
}
