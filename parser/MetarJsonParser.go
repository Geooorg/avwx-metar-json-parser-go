package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const URL_TEMPLATE = "http://avwx.rest/api/metar/%s?token=%s"
const HTTP_TIMEOUT = 10

type JsonStruct struct {
	Raw           string `json:"raw"`
	WindDirection struct {
		Value int `json:"value"`
	} `json:"wind_direction"`

	WindSpeed struct {
		Value int `json:"value"`
	} `json:"wind_speed"`

	Temperature struct {
		Value int `json:"value"`
	} `json:"temperature"`

	TemperatureUnit struct {
		Value string `json:"temperature"`
	} `json:"units"`
}

func ConvertToMetarData(metarJson JsonStruct) MetarData {
	return MetarData{
		Raw:            metarJson.Raw,
		WindDirection:  metarJson.WindDirection.Value,
		WindSpeed:      metarJson.WindSpeed.Value,
		Temperatur:     metarJson.Temperature.Value,
		TemperaturUnit: metarJson.TemperatureUnit.Value,
	}
}

func ReadJsonFromWebservice() ([]byte, error) {

	paramaterizedUrl := GetParameterizedUrl()

	log.Printf("DEBUG: Reading JSON from %s", URL_TEMPLATE)

	client := http.Client{
		Timeout: HTTP_TIMEOUT * time.Second,
	}
	response, err := client.Get(paramaterizedUrl)
	if err != nil {
		log.Printf("WARN: Reading data from web service failed: %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var data bytes.Buffer
	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func GetParameterizedUrl() string {

	token := os.Getenv("AVWX_TOKEN")
	airportCode := os.Getenv("AVWX_AIRPORT")

	return fmt.Sprintf(URL_TEMPLATE, airportCode, token)
}
