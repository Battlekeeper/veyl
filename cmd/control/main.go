package main

import (
	"net/http"
	"sync"

	"github.com/Battlekeeper/veyl/cmd/control/router"
	"github.com/Battlekeeper/veyl/internal/database"
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Relays      = make(map[string]types.RelayClient)
	RelaysMutex = &sync.Mutex{} // Ensure thread-safe access to Relays
)

func main() {
	database.StartClient()

	// r.GET("/register", func(c *gin.Context) {
	// 	// get public key from query parameters
	// 	publicKey := c.Query("public_key")
	// 	if publicKey == "" {
	// 		c.JSON(400, gin.H{"error": "public_key is required"})
	// 		return
	// 	}
	// 	relayId := c.Query("relay_id")
	// 	if relayId == "" {
	// 		c.JSON(400, gin.H{"error": "relay_id is required"})
	// 		return
	// 	}
	// 	port := c.Query("port")
	// 	if port == "" {
	// 		c.JSON(400, gin.H{"error": "port is required"})
	// 		return
	// 	}
	// 	ip := c.Query("ip")
	// 	if ip == "" {
	// 		c.JSON(400, gin.H{"error": "ip is required"})
	// 		return
	// 	}
	// 	log.Printf("Registering relay %s with public key %s, port %s, ip %s", relayId, publicKey, port, ip)
	// 	portInt, err := strconv.Atoi(port)
	// 	if err != nil {
	// 		c.JSON(400, gin.H{"error": "port must be an integer"})
	// 		return
	// 	}
	// 	connection := types.RelayConnection{
	// 		RelayID:   relayId,
	// 		PublicKey: publicKey,
	// 		IP:        ip,
	// 		Port:      portInt,
	// 	}
	// 	relay, ok := Relays[relayId]
	// 	if !ok {
	// 		c.JSON(404, gin.H{"error": "relay not connected"})
	// 		return
	// 	}
	// 	relay.Connection.WriteJSON(connection)
	// 	c.JSON(200, relay.Auth)
	// })
	// r.GET("/relay", func(c *gin.Context) {
	// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// 	if err != nil {
	// 		log.Println("WebSocket upgrade failed:", err)
	// 		return
	// 	}
	// 	defer conn.Close()

	// 	auth := types.RelayAuth{}
	// 	err = conn.ReadJSON(&auth)
	// 	if err != nil {
	// 		log.Println("read error:", err)
	// 		return
	// 	}
	// 	RelayConnection := types.RelayClient{
	// 		Auth:       auth,
	// 		Connection: conn,
	// 	}
	// 	RelaysMutex.Lock()
	// 	Relays[auth.RelayID] = RelayConnection
	// 	RelaysMutex.Unlock()
	// 	log.Printf("Relay %s connected with public key %s", auth.RelayID, auth.PublicKey)

	// 	defer func() {
	// 		RelaysMutex.Lock()
	// 		delete(Relays, auth.RelayID)
	// 		RelaysMutex.Unlock()
	// 		log.Printf("Relay %s disconnected", auth.RelayID)
	// 	}()
	// 	for {
	// 		_, _, err := conn.ReadMessage()
	// 		if err != nil {
	// 			break
	// 		}
	// 	}
	// })

	router.Initialize()
}
