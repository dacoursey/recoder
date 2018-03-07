package main

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type retort struct {
	ID      string `json:"id"`
	Base64f string `json:"base64f"`
	Base64b string `json:"base64b"`
	Base32f string `json:"base32f"`
	Base32b string `json:"base32b"`
	Hexf    string `json:"hexf"`
	Hexb    string `json:"hexb"`
	Binf    string `json:"binf"`
	Binb    string `json:"binb"`
}

func (a *App) recode(w http.ResponseWriter, r *http.Request) {
	i := r.FormValue("input")

	// If no input value, send back error.

	uuid, err := newUUID()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Base64
	b64f := base64.StdEncoding.EncodeToString([]byte(i))
	b64b, err := base64.StdEncoding.DecodeString(i)

	// Base32
	b32f := base32.StdEncoding.EncodeToString([]byte(i))
	b32b, err := base32.StdEncoding.DecodeString(i)

	// Hexidecimal
	hxf := hex.EncodeToString([]byte(i))
	hxb, err := hex.DecodeString(i)

	// Binary
	buff := new(bytes.Buffer)
	err = binary.Write(buff, binary.LittleEndian, i)
	binf := buff.String()

	var binb string
	b := []byte(i)
	bufb := bytes.NewReader(b)
	binary.Read(bufb, binary.LittleEndian, &binb)

	respondWithJSON(w, http.StatusOK, retort{uuid,
		string(b64f), string(b64b),
		string(b32f), string(b32b),
		string(hxf), string(hxb),
		binf, binb})
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
