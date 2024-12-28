package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Returns responce body from an API call. JSON expected.
func GetRespBody(apiCall string) []byte {
	resp, err := http.Get(apiCall)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return body
}

// Converts JSON from raw responce into GO type
func TypefyResp(data []byte) map[string]interface{} {

	// Chew raw data: init variable with arbitrary data structure
	var digested map[string]interface{}
	if err := json.Unmarshal(data, &digested); err != nil {
		panic(err)
	}

	return digested
}
