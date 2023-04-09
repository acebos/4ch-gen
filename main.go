package main

import (
	"encoding/json"
	"os"
)

var (
	BOARDS = []string{"g", "fit", "biz", "vg", "vr", "ck", }
)

func main() {
	mappings := make(map[string]interface{})

	for _, board := range BOARDS {
		catalog, err := GetBoardCatalog(board)
		if err != nil {
			panic(err)
		}

		generals := BuildGenerals(board, catalog)
		mappings[board] = generals
	}

	jsonPayload, err := json.MarshalIndent(mappings, "", "    ")
	if err != nil {
		panic(err)
	}

	os.WriteFile("output/mappings.json", jsonPayload, 0644)
}
