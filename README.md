# avwx-metar-json-parser-go

This little library helps you to retrieve METAR weather data kindly provided by http://avwx.rest. 

At the moment, only the METAR raw string, wind direction and speed as well as temperatur are mapped to the 'Metar.go' domain object

## Environment parameters to pass

Export the following environment parameters:
 - AVWX_AIRPORT (e.g. EDDF for the airport Frankfurt/Main)
 - AVWX_TOKEN for your API token. This can be requested at avwx.rest