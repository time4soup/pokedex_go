package poke_api_client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// makes get request to given url and returns json body data
func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return nil, fmt.Errorf("Not Found")
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}
