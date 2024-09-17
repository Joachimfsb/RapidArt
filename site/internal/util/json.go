package util

import (
	"encoding/json"
	"io"
)

// decodeBody decodes the body of a given response into a given value (content).
func JsonDecode(body io.ReadCloser, content any) error {
	decoder := json.NewDecoder(body) // Initialize the decoder
	// Decode the body into the data-type
	if err1 := decoder.Decode(content); err1 != nil {
		return err1
	}
	err2 := body.Close() // Closing the body.
	if err2 != nil {
		return err2
	}
	// Everything is OK
	return nil
}

func JsonEncode(content any) ([]byte, error) {
	return json.MarshalIndent(content, "", "  ")
}
