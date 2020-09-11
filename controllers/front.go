package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers :
// Register controller using Handler interface
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// encodeResponseAsJSON :
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
