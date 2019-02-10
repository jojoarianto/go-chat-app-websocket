package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

    "github.com/jojoarianto/go-chat-app-websocket/domain"
	"github.com/julienschmidt/httprouter"
)

// sentCollectionMsg collection of sent out message
var sentCollectionMsg []domain.Message

// Run start server
func Run(port int) error {
	// Start websocket listening for incoming chat messages
	go handleMessages()
	
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/sent", getSentMessage)
	r.POST("/sent", sentMessage)
	r.GET("/ws", handleConnections)

	return r
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	respondWithJson(w, http.StatusOK, "WELCOME TO GO CHAT API")
}

func getSentMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	respondWithJson(w, http.StatusOK, sentCollectionMsg)
}

func sentMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type userRequest struct {
		ContentMessage string `json:"content_message"`
	}
	var ureq userRequest // catch user request

	// convert json to object
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ureq); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// Send received message to the broadcast channel
	broadcast <- domain.NewMessage(ureq.ContentMessage)

	msg := domain.NewMessage(ureq.ContentMessage)
	sentCollectionMsg = append(sentCollectionMsg, msg)
	respondWithJson(w, http.StatusOK, msg)
}
