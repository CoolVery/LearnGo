package main

import (
	"io"
	"strings"

	. "github.com/CoolVery/LearnGo.git/classes"
)

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

func Contains(r io.Reader, seq []byte) (bool, error) {
	data := make([]byte, len(seq))
	readByte, errRead := r.Read(data)
	if errRead != nil && errRead != io.EOF {
		return false, errRead
	}
	stringSeq := string(seq)
	stringRead := string(data[readByte:])
	if strings.Contains(stringRead, stringSeq) {
		return true, nil
	} else {
		return false, nil
	}
}

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, n)
	readByte, errRead := r.Read(data)
	if errRead != nil && errRead != io.EOF {
		return errRead
	}
	_, errWrite := w.Write(data[:readByte])
	if errWrite != nil {
		return errWrite
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
	_, err := writer.Write([]byte(text))
	if err != nil {
		return err
	}
	return nil
		
}

func main() {

}
