package cmd

import (
	"go-mt2/config"
	"go-mt2/internal/packets/handlers"
	"go-mt2/internal/packets/in"
	"go-mt2/pkg/auth"
	"go-mt2/pkg/db"
	"log"
	"net/http"

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

// Add packet handling method to App
func (app *App) HandlePacket(packet in.ClientPacket, conn *websocket.Conn) error {
	switch p := packet.(type) {
	case *in.LoginPacket:
		return app.handleLogin(p, conn)
	// Add other packet types here
	default:
		log.Printf("Unhandled packet type: %T", packet)
		return nil
	}
}

// Example login handler
func (app *App) handleLogin(packet *in.LoginPacket, conn *websocket.Conn) error {
	log.Printf("Login attempt: username=%s", packet.Username)

	// Use your auth service here
	// result := app.Auth.ValidateLogin(packet.Username, packet.Password)

	// Send response back to client
	// response := CreateLoginResponse(result)
	// return conn.WriteMessage(websocket.BinaryMessage, response)

	return nil
}
