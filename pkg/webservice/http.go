package webservice

import (
	"encoding/json"
	"io"
	"net/http"
)

func RespondWithStatus(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func ReadFullBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}
