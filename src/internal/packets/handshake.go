package packets

type Handshake struct {
	Id    int
	Time  int
	Delta int
}

func (h *Handshake) GetId() int {
	return h.Id
}

func (h *Handshake) GetTime() int {
	return h.Time
}

func (h *Handshake) GetDelta() int {
	return h.Delta
}
