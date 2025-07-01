package packets

type PacketBidirectional struct {
	writer *Writer
}

type Writer struct {
	buffer []byte
	offset int
}

func (w *Writer) WriteUInt8(data uint8) *Writer {
	w.buffer = append(w.buffer, data)
	return w
}
func (w *Writer) WriteUInt16(data uint16) *Writer {
	w.buffer = append(w.buffer, byte(data>>8), byte(data))
	return w
}
func (w *Writer) WriteUInt32(data uint32) *Writer {
	w.buffer = append(w.buffer, byte(data>>24), byte(data>>16), byte(data>>8), byte(data))
	return w
}
func (w *Writer) WriteUInt64(data uint64) *Writer {
	w.buffer = append(w.buffer, byte(data>>56), byte(data>>48), byte(data>>40), byte(data>>32), byte(data>>24), byte(data>>16), byte(data>>8), byte(data))
	return w
}
func (w *Writer) WriteBytes(data []byte) *Writer {
	w.buffer = append(w.buffer, data...)
	return w
}
func (w *Writer) WriteString(data string) *Writer {
	w.buffer = append(w.buffer, []byte(data)...)
	return w
}

func (w *Writer) GetBuffer() []byte {
	return w.buffer
}

func (w *Writer) Reset() {
	w.buffer = w.buffer[:0]
	w.offset = 0
}
