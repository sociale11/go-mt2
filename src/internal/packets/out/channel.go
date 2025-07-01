package out

import "go-mt2/internal/packets"

/**
* @type: out
* @header: 0x79
* @size: 2 bytes
* @description: Used to send the number of channel.
 */
type ChannelPacket struct {
	Header  byte
	Channel byte
}

func (p *ChannelPacket) getHeader() byte {
	return 0x79
}

func (p *ChannelPacket) Parse(reader *packets.BufferReader) error {
	var err error
	p.Header, err = reader.ReadByte()
	if err != nil || p.Header != p.getHeader() {
		return err
	}

	p.Channel, err = reader.ReadByte()
	if err != nil {
		return err
	}

	return nil
}
