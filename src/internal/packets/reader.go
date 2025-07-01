package packets

import (
	"encoding/binary"
	"errors"
)

var _ PacketReader = (*BufferReader)(nil)

type BufferReader struct {
	buffer []byte
	offset int
}

type PacketReader interface {
	SetBuffer(buffer []byte, skipHeader bool)
	ReadByte() (byte, error)
	ReadUInt8() (uint8, error)
	ReadUInt16() (uint16, error)
	ReadUInt32() (uint32, error)
	ReadUInt64() (uint64, error)
	ReadString(length int) (string, error)
}

func NewBufferReader(data []byte, offset int) *BufferReader {
	return &BufferReader{buffer: data, offset: offset}
}

func (r *BufferReader) SetBuffer(buffer []byte, skipHeader bool) {
	r.Reset()
	r.buffer = buffer

	if skipHeader {
		r.offset = 1
	}
}

func (r *BufferReader) ReadByte() (byte, error) {
	res, err := r.ReadField(1)
	if err != nil {
		return 0, err
	}
	return res[0], nil
}

func (r *BufferReader) ReadUInt8() (uint8, error) {
	res, err := r.ReadField(1)
	if err != nil {
		return 0, err
	}
	return uint8(res[0]), nil
}

func (r *BufferReader) ReadUInt16() (uint16, error) {
	res, err := r.ReadField(2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(res), nil
}

func (r *BufferReader) ReadUInt32() (uint32, error) {
	res, err := r.ReadField(4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(res), nil
}

func (r *BufferReader) ReadUInt64() (uint64, error) {
	res, err := r.ReadField(8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(res), nil
}

func (r *BufferReader) ReadString(length int) (string, error) {
	res, err := r.ReadField(length)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (r *BufferReader) GetBuffer() []byte {
	return r.buffer
}

func (r *BufferReader) Reset() {
	r.buffer = r.buffer[:0]
	r.offset = 0
}

func (r *BufferReader) checkBounds(size int) error {
	if r.offset+size > len(r.buffer) {
		return errors.New("BufferReader: Reading past the end of the buffer")
	}
	return nil
}

func (r *BufferReader) ReadField(size int) ([]byte, error) {
	err := r.checkBounds(size)
	if err != nil {
		return nil, err
	}
	res := r.buffer[r.offset : r.offset+size]
	r.offset += size
	return res, nil
}
