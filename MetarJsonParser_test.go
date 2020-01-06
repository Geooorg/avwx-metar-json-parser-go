package main

import (
	"./parser"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestUrlIsParameterizedCorrectly(t *testing.T) {
	os.Setenv("AVWX_TOKEN", "abcTOKEN123")
	os.Setenv("AVWX_AIRPORT", "EDDF")

	url := parser.GetParameterizedUrl()
	assert.Equal(t, url, "http://avwx.rest/api/metar/EDDF?token=abcTOKEN123", "URL matches")
}

func TestDataCanBeConverted(t *testing.T) {
	jsonFile, err := os.Open("data/test.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var metarDataJson parser.JsonStruct
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &metarDataJson)

	metar := parser.ConvertToMetarData(metarDataJson)

	fmt.Println(metar.Raw)

	assert.Equal(t, metar.Raw, "EDDF 062050Z 19008KT 9999 FEW021 04/01 Q1019 NOSIG")
	assert.Equal(t, metar.WindDirection, 190)
	assert.Equal(t, metar.WindSpeed, 8)
	assert.Equal(t, metar.Temperatur, 4)
	assert.Equal(t, metar.TemperaturUnit, "C")
}
