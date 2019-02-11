package interfaces

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/jojoarianto/go-chat-app-websocket/domain"
)

// initialize the upgrader
var upgrader = websocket.Upgrader{}
// collection of connected clients
var clients = make(map[*websocket.Conn] bool) 
// broadcast channel
var broadcast = make(chan domain.Message)

// read broadcast channel
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		
		// Send it out to every client that is currently connected
		for client := range clients {
			log.Printf("msg from client: %v", msg.ContentMessage)
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// init handler for websocket
func handleConnections(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	
	// Save new client
	clients[ws] = true
	// Set Keep alive (60 minutes)
	keepAlive(ws, (60 * time.Minute))
	
	for {
		var msg domain.Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		// Send received message to the broadcast channel
		broadcast <- domain.NewMessage(msg.ContentMessage)
	}
}

// keepAlive function to set timeout of websocket
func keepAlive(c *websocket.Conn, timeout time.Duration) {
    lastResponse := time.Now()
    c.SetPongHandler(func(msg string) error {
       lastResponse = time.Now()
       return nil
	})

	go func() {
    	for {
			err := c.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil { return }

			time.Sleep(timeout/2)
			if(time.Now().Sub(lastResponse) > timeout) {
				c.Close()
				return
			}
    	}
  	}()
}