package main

import (
	"github.com/Geooorg/avwx-metar-json-parser-go/parser"
	"log"
)

func main() {
	metar, e := parser.GetMetarData()
	if e != nil {
		log.Println("WARN: Could not read data from web service")
	}
	log.Printf("METAR %s", metar)
}
