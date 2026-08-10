package classes

import "strings"

type UpperWriter struct {
	UpperString string
}

func (upperWriter *UpperWriter) Write(text []byte) (int, error) {
	stringText := strings.ToUpper(string(text))

	upperWriter.UpperString = strings.ToUpper(text)
	return nil
}