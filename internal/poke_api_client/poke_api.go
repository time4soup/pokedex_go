package poke_api_client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func PokeApiGet(url string) MapResponse {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	mapRes := MapResponse{}
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		log.Fatal(err)
	}

	return mapRes
}
