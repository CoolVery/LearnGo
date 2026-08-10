package main

import (
	"io"
	. "github.com/CoolVery/LearnGo.git/classes"
)

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 1024)
	readByte, err := r.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}
	
	return string(data[:readByte]), nil
}

func WriteData(text string, writer UpperWriter) error {
	err := writer.Write(text)
	if err != nil {
		return err
	}
	return nil
		
}

func main() {

}
