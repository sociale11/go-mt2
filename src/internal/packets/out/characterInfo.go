package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x20
* @size: 329 bytes
* @description: Is used to send the characters list to client select screen (we need to repeat the characterInfo 4x, guildIds 4x, guildNames 4x).
 */
type CharactersInfoPacket struct {
	Header      byte
	Id          uint32
	Name        string
	PlayerClass byte
	Level       byte
	Playtime    uint32
	St          byte
	Ht          byte
	Dx          byte
	Iq          byte
	BodyPart    uint16
	NameChange  byte
	HairPart    uint16
	Unknown1    byte //0
	PositionX   uint32
	PositionY   uint32
	Ip          uint32
	Port        uint32
	SkillGroup  byte
	GuildId     uint32
	GuildName   string
	Unknown2    byte //0
	Unknown3    byte //0
}

func (p *CharactersInfoPacket) getHeader() byte {
	return 0x20
}

func (p *CharactersInfoPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}
	p.Id, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.Name, err = reader.ReadString(25)
	if err != nil {
		return err
	}
	p.PlayerClass, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Level, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Playtime, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.St, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Ht, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Dx, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Iq, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.BodyPart, err = reader.ReadUInt16()
	if err != nil {
		return err
	}
	p.NameChange, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.HairPart, err = reader.ReadUInt16()
	if err != nil {
		return err
	}
	p.Unknown1, err = reader.ReadByte()
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
	p.Ip, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.Port, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.SkillGroup, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.GuildId, err = reader.ReadUInt32()
	if err != nil {
		return err
	}
	p.GuildName, err = reader.ReadString(13)
	if err != nil {
		return err
	}
	p.Unknown2, err = reader.ReadByte()
	if err != nil {
		return err
	}
	p.Unknown3, err = reader.ReadByte()
	if err != nil {
		return err
	}

	return nil
}
