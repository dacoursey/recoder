package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type retort struct {
	ID          string `json:"id"`
	Base64f     string `json:"base64f"`
	Base64b     string `json:"base64b"`
	Base32      string `json:"base32"`
	Hexidecimal string `json:"hex"`
}

func (a *App) recode(w http.ResponseWriter, r *http.Request) {
	i := r.FormValue("input")

	// If no input value, send back error.

	uuid, err := newUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	b64f := base64.StdEncoding.EncodeToString([]byte(i))
	b64b, err := base64.StdEncoding.DecodeString(i)
	b32, err := base32.StdEncoding.DecodeString(i)
	hx, err := hex.DecodeString(i)

	respondWithJSON(w, http.StatusOK, retort{uuid, string(b64f), string(b64b), string(b32), string(hx)})
}

/////
// Handlers Section
/////

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
