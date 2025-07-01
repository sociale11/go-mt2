package in

import "go-mt2/internal/packets"

type LoginPacket struct {
	Header   ClientPacketType
	Username string
	Password string
}

func (p *LoginPacket) GetHeader() byte {
	return byte(p.Header)
}

func (p *LoginPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Username, err = reader.ReadString(30)
	if err != nil {
		return err
	}

	p.Password, err = reader.ReadString(16)
	if err != nil {
		return err
	}
	return nil
}

func (p *LoginPacket) GetType() byte {
	return byte(HEADER_CG_LOGIN)
}
