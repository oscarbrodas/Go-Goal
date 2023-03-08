package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DecodeJSONRequest(dst interface{}, src io.Reader, w http.ResponseWriter) {
	dec := json.NewDecoder(src)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request JSON (%v).", err)
	}
}
