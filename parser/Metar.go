package parser

import (
	"bytes"
	"fmt"
)

type MetarData struct {
	Raw            string
	WindDirection  int
	WindSpeed      int
	Temperatur     int
	TemperaturUnit string
}

func (it MetarData) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(", Raw " + it.Raw)
	buffer.WriteString(", WindDirection " + fmt.Sprint(it.WindDirection))
	buffer.WriteString(", WindSpeed " + fmt.Sprintf("%f", it.WindSpeed))
	buffer.WriteString(", Temperatur " + fmt.Sprintf("%f", it.Temperatur) + it.TemperaturUnit)

	return buffer.String()
}
