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
	buffer.WriteString(", WindDirection " + fmt.Sprintf("%d", it.WindDirection))
	buffer.WriteString(", WindSpeed " + fmt.Sprintf("%d", it.WindSpeed))
	buffer.WriteString(", Temperatur " + fmt.Sprintf("%d ", it.Temperatur) + "°" + it.TemperaturUnit)

	return buffer.String()
}
