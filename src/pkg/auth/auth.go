package auth

import "golang.org/x/net/websocket"

type AuthService struct {
	Service map[*websocket.Conn]string
}
