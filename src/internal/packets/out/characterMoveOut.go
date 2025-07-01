package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x03
* @size: 25 bytes
* @description: Is used to replicate the movement of a character (player, mobs) to other nearby players.
 */
type CharacterMoveOutPacket struct {
	Header       byte
	MovementType byte
	Arg          byte
	Rotation     byte
	Vid          uint32
	PositionX    uint32
	PositionY    uint32
	Time         uint32
	Duration     uint32
	Unknown      byte
}

func (p *CharacterMoveOutPacket) getHeader() byte {
	return 0x03
}

func (p *CharacterMoveOutPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}

	p.MovementType, err = reader.ReadByte()
	if err != nil {
		return err
	}

	p.Arg, err = reader.ReadByte()
	if err != nil {
		return err
	}

	p.Rotation, err = reader.ReadByte()
	if err != nil {
		return err
	}

	p.Vid, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.PositionX, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.PositionY, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.Time, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.Duration, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.Unknown, err = reader.ReadByte()
	if err != nil {
		return err
	}
	return nil
}
