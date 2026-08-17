package main

import (
	"io"
	"bytes"
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
	buf := make([]byte, 4094)
	window := make([]byte, 0, len(seq)*2)

	for {
		n, err := r.Read(buf)
		if n > 0 {
			window = append(window, buf[:n]...)

			if bytes.Contains(window, seq) {
				return true, nil
			}

			if len(window) > len(buf) {
				window = window[len(window)-len(seq):]
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
	}

	return false, nil
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
