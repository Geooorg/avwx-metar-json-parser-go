package main

import (
	"./parser"
	"encoding/json"
	"log"
)

func main() {

	jsonStr, e := parser.ReadJsonFromWebservice()
	if e != nil {
		log.Fatal("Could not read data from web service")
	}

	var metarDataJson parser.JsonStruct
	json.Unmarshal(jsonStr, &metarDataJson)

	metar := parser.ConvertToMetarData(metarDataJson)
	log.Printf("METAR %s", metar)
}
