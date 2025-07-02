package in

import (
	"errors"
	"go-mt2/internal/packets"

	"go-mt2/pkg/db"
	"go-mt2/pkg/db/models"
	"log"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

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

// Example login handler
func HandleLogin(packet *LoginPacket, conn *websocket.Conn) error {
	log.Printf("Login attempt: username=%s password=%s", packet.Username, packet.Password)
	var user models.Account
	db.DB.AuthDB.Where("username = ?", packet.Username).First(user)

	if user.Password != HashPassword(packet.Password) {
		return errors.New("Invalid username or password")
	} else {
		log.Println("Login successful")
	}

	// Use your auth service here
	// result := app.Auth.ValidateLogin(packet.Username, packet.Password)

	// Send response back to client
	// response := CreateLoginResponse(result)
	// return conn.WriteMessage(websocket.BinaryMessage, response)

	return nil
}

func HashPassword(password string) string {
	cost := 12
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Printf("Error generating password hash: %v", err)
		return ""
	}
	return string(bytes)
}

func ValidateLogin(username, password string) bool {
	var user models.Account
	db.DB.AuthDB.Where("username = ?", username).First(user)

	if user.Password != HashPassword(password) {
		return false
	}
	return true
}
