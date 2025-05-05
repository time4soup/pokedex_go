package poke_api_client

// type MapItem struct {
// 	Name string `json:"name"`
// 	Url  string `json:"url"`
// }

type MapResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

// type locationArea struct {
// 	Id int `json:"version_details"`
// 	Name string `json:"name"`
// 	GameIndex int `json:"game_index"`
// 	EncounterMethodRate encounterMethodRate `json:"encounter_method_rates"`
// 	Location location `json:"location"`
// 	names
// 	pokemon_encounters
// }

// type encounterMethodRate struct {
// 	encounter_method
// 	version_details
// }

// type location struct {
// 	Id int `json:"id"`
// 	Name string `json:"name"`
// 	region
// 	names
// 	game_indices
// 	Areas []locationArea `json:"areas"`
// }
