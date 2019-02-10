package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

    "github.com/jojoarianto/p1/domain"
	"github.com/julienschmidt/httprouter"
)

// sentCollectionMsg is collection of sent message
var sentCollectionMsg []domain.Message

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/sent", getSentMessage)
	r.POST("/sent", sentMessage)

	return r
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	respondWithJson(w, http.StatusOK, "GO CHAT API")
}

func getSentMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	respondWithJson(w, http.StatusOK, sentCollectionMsg)
}

func sentMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var ureq userRequest // catch user request 

	// convert json to object
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ureq); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// create msg object by with content msg
	msg := domain.NewMessage(ureq.ContentMessage)

	sentCollectionMsg = append(sentCollectionMsg, msg)
	respondWithJson(w, http.StatusOK, msg)
}
