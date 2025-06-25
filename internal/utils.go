package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func NewDecoder(r io.Reader) *json.Decoder {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	return dec
}

func NewEncoder(w io.Writer) *json.Encoder {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	return enc
}

// Writes obj into w with status header
func WriteJSON(w http.ResponseWriter, status int, obj any) {
	enc := NewEncoder(w)

	w.WriteHeader(status)
	if err := enc.Encode(obj); err != nil {
		fmt.Printf("Encode failed: %v\n", err)
		slog.Error(fmt.Sprintf("writing response: %v", err))
	}
}

// Reads JSON from r.Body to dest
func ReadJSON(r *http.Request, dest any) error {
	dec := NewDecoder(r.Body)

	return dec.Decode(dest)
}

