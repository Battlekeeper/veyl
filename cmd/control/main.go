package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Battlekeeper/veyl/internal/database"
	"github.com/Battlekeeper/veyl/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	net := types.CreateNetwork("Test Network")
	id := primitive.NewObjectID()
	net.AddRelay(types.Relay{
		Id:        id, // Example ObjectID
		PublicKey: "examplePublicKey",
		Name:      "Test Relay",
		NetworkId: net.Id,
	})

	relay, err := types.GetRelayById(id)
	if err != nil {
		log.Fatalf("Failed to get relay by ID: %v", err)
	}
	fmt.Printf("Relay ID: %s, Public Key: %s, Name: %s\n", relay.Id.Hex(), relay.PublicKey, relay.Name)

	return

	log.Printf("Created network: %s with ID: %s", net.Name, net.Id.Hex())
	net.AddRelay(types.Relay{
		Id:        primitive.NewObjectID(), // Example ObjectID
		PublicKey: "examplePublicKey",
		Name:      "Test Relay",
		NetworkId: net.Id,
	})
	time.Sleep(5 * time.Second) // Simulate some delay for the network setup
	net.Update()

	r := gin.Default()
	r.GET("/register", func(c *gin.Context) {
		// get public key from query parameters
		publicKey := c.Query("public_key")
		if publicKey == "" {
			c.JSON(400, gin.H{"error": "public_key is required"})
			return
		}
		relayId := c.Query("relay_id")
		if relayId == "" {
			c.JSON(400, gin.H{"error": "relay_id is required"})
			return
		}
		port := c.Query("port")
		if port == "" {
			c.JSON(400, gin.H{"error": "port is required"})
			return
		}
		ip := c.Query("ip")
		if ip == "" {
			c.JSON(400, gin.H{"error": "ip is required"})
			return
		}
		log.Printf("Registering relay %s with public key %s, port %s, ip %s", relayId, publicKey, port, ip)
		portInt, err := strconv.Atoi(port)
		if err != nil {
			c.JSON(400, gin.H{"error": "port must be an integer"})
			return
		}
		connection := types.RelayConnection{
			RelayID:   relayId,
			PublicKey: publicKey,
			IP:        ip,
			Port:      portInt,
		}
		relay, ok := Relays[relayId]
		if !ok {
			c.JSON(404, gin.H{"error": "relay not connected"})
			return
		}
		relay.Connection.WriteJSON(connection)
		c.JSON(200, relay.Auth)
	})
	r.GET("/relay", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			return
		}
		defer conn.Close()

		auth := types.RelayAuth{}
		err = conn.ReadJSON(&auth)
		if err != nil {
			log.Println("read error:", err)
			return
		}
		RelayConnection := types.RelayClient{
			Auth:       auth,
			Connection: conn,
		}
		RelaysMutex.Lock()
		Relays[auth.RelayID] = RelayConnection
		RelaysMutex.Unlock()
		log.Printf("Relay %s connected with public key %s", auth.RelayID, auth.PublicKey)

		defer func() {
			RelaysMutex.Lock()
			delete(Relays, auth.RelayID)
			RelaysMutex.Unlock()
			log.Printf("Relay %s disconnected", auth.RelayID)
		}()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
