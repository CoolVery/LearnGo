package classes

import (
	_ "strings"
)

type UpperReaderWriter struct {
	UpperString (string)
}

func (uprw UpperReaderWriter) Read([]byte) {

}

func (uprw UpperReaderWriter) Write(p []byte) int {
	
	return 0
}
