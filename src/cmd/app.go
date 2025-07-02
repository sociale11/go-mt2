package cmd

import (
	"encoding/binary"
	"go-mt2/config"
	"go-mt2/internal/packets/handlers"
	"go-mt2/internal/packets/in"
	"go-mt2/pkg/auth"
	"go-mt2/pkg/db"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type App struct {
	// Logger Logger
	Auth   auth.AuthService
	DB     db.DBService
	Config config.Config
}

func Execute() *App {
	// Initialize services
	dbService := db.DBService{
		AuthDB: db.NewAuthDB(),
		GameDB: db.NewGameDB(),
	}

	dbService.SeedAuth()

	app := &App{
		// Logger: Logger{},
		Auth:   auth.AuthService{},
		DB:     dbService, // Use the actual dbService, not empty struct
		Config: config.Settings,
	}

	// Set up HTTP server with WebSocket endpoint
	http.HandleFunc("/ws", app.HandleWebSocket)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	return app // This won't be reached due to ListenAndServe, but needed for signature
}

// Add this method to your App struct
func (app *App) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := handlers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer c.Close()

	log.Println("Client connected")
	log.Println("Creating handshake packet")

	c.WriteMessage(websocket.BinaryMessage, app.CreateHandshakePacket())

	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		// Parse the packet
		packet, err := handlers.CreatePacketFromData(data)
		if err != nil {
			log.Printf("Packet parse error: %v", err)
			continue
		}

		// Handle the packet with your services
		err = app.HandlePacket(packet, c)
		if err != nil {
			log.Printf("Packet handling error: %v", err)
		}
	}
}

func (app *App) HandlePacket(packet in.ClientPacket, conn *websocket.Conn) error {
	log.Println(conn.RemoteAddr())
	switch p := packet.(type) {
	case *in.LoginPacket:
		return in.HandleLogin(p, conn)
	case *in.HandshakePacket:
		return in.HandleHandshake(p, conn)
	default:
		log.Printf("Unhandled packet type: %T", packet)
		return nil
	}
}

func (app *App) CreateHandshakePacket() []byte {
	data := make([]byte, 1+4+4+4)

	data[0] = byte(in.HEADER_CG_HANDSHAKE)
	dwHandshake := rand.Uint32()
	dwTime := uint32(time.Now().Unix())
	lDelta := 0

	binary.LittleEndian.PutUint32(data[1:5], uint32(dwHandshake))
	binary.LittleEndian.PutUint32(data[5:9], uint32(dwTime))
	binary.LittleEndian.PutUint32(data[9:13], uint32(lDelta))

	return data
}
