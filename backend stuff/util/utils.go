package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func Base64toFile(base64Encoding string) ([]byte, error) {
	if !strings.HasPrefix(base64Encoding, "data:image/png;base64,") {
		return []byte{}, fmt.Errorf("Base64 encoding does not have correct prefix")
	}

	input := strings.TrimPrefix(base64Encoding, "data:image/png;base64,")
	return base64.StdEncoding.DecodeString(input)
}

func FiletoBase64(fileBytes []byte) string {
	return base64.RawStdEncoding.EncodeToString(fileBytes)
}
