package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
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

func WriteJSON(w http.ResponseWriter, status int, obj any) {
	enc := NewEncoder(w)

	w.WriteHeader(status)
	if err := enc.Encode(obj); err != nil {
		fmt.Printf("Encode failed: %v\n", err)
		slog.Error(fmt.Sprintf("writing response: %v", err))
	}
}

func ReadJSON(r *http.Request, dest any) error {
	dec := NewDecoder(r.Body)

	return dec.Decode(dest)
}

func FormatTime(d time.Duration) string {
	if d == 0 {
		return "0s"
	}

	negative := d < 0
	if negative {
		d = -d
	}

	var result string

	if d >= time.Minute {
		minutes := d / time.Minute
		d %= time.Minute
		seconds := d / time.Second
		d %= time.Second

		if seconds > 0 {
			result = fmt.Sprintf("%dm %ds", minutes, seconds)
		} else {
			result = fmt.Sprintf("%dm", minutes)
		}

		if d >= time.Millisecond {
			ms := d / time.Millisecond
			result += fmt.Sprintf(" %dms", ms)
		}
	} else if d >= time.Second {
		seconds := d / time.Second
		d %= time.Second

		if d >= time.Millisecond {
			ms := d / time.Millisecond
			result = fmt.Sprintf("%ds %dms", seconds, ms)
		} else {
			result = fmt.Sprintf("%ds", seconds)
		}
	} else if d >= time.Millisecond {
		result = fmt.Sprintf("%dms", d/time.Millisecond)
	} else if d >= time.Microsecond {
		result = fmt.Sprintf("%dµs", d/time.Microsecond)
	} else {
		result = fmt.Sprintf("%dns", d)
	}

	if negative {
		result = "-" + result
	}

	return result
}
