package classes

import (
	. "strings"
	"unicode"
)

type UpperReaderWriter struct {
	UpperString (string)
}

func (uprw *UpperReaderWriter) Read() []byte{
	return []byte(uprw.UpperString)
}

func (uprw *UpperReaderWriter) Write(p []byte) int {
	resultWriteByte := 0
	runes := []rune(string(p))
	uprw.UpperString = ""
	for _, value := range runes {
		uprw.UpperString += ToUpper(string(value))
		if unicode.Is(unicode.Cyrillic, value) {
			resultWriteByte += 2
		} else {
		resultWriteByte++
		}
	}
	return resultWriteByte
}
