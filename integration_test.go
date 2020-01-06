package main

import (
	"encoding/json"
	"github.com/Geooorg/avwx-metar-json-parser-go/parser"
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
