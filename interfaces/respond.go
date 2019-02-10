package interfaces

import (
	"encoding/json"
	"net/http"
)

// userRequest format object
type userRequest struct {
	ContentMessage string `json:"content_message"`
}

type userRespond struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// method to print error output for http respon
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// method to print output for http respon
// parameter [w (Http.RestponWriter), http.statuscode, payload/data/msg]
// payload is data credential which will be trans to other part
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}