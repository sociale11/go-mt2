package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x88
* @size: 54 bytes
* @description: Represents the basic information about the character.
 */
type CharacterDetailsPacket struct {
	Header     byte
	Vid        uint32
	PlayerName string
	Parts      uint32
	EmpireId   byte
	GuildId    uint32
	Level      uint32
	RankPoints uint16
	PkMode     byte
	MountId    uint32
}

func (p *CharacterDetailsPacket) getHeader() byte {
	return 0x88
}

func (p *CharacterDetailsPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}

	p.Vid, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.PlayerName, err = reader.ReadString(32)
	if err != nil {
		return err
	}

	p.Parts, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.EmpireId, err = reader.ReadByte()
	if err != nil {
		return err
	}

	p.GuildId, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.Level, err = reader.ReadUInt32()
	if err != nil {
		return err
	}

	p.RankPoints, err = reader.ReadUInt16()
	if err != nil {
		return err
	}

	p.PkMode, err = reader.ReadByte()
	if err != nil {
		return err
	}

	p.MountId, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	return nil
}
