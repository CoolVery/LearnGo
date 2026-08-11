package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
	"time"
	"errors"
)

func ReadContent(filename string) string {
	buffer := make([]byte, 100)
	resultString := ""
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return ""
	}
	for {
		readByte, err := file.Read(buffer)
		if err == io.EOF {
			resultString += string(buffer)[:readByte]
			return resultString
		}
		if err != nil {
			return ""
		}
		resultString += string(buffer)[:readByte]
	}
}

func readContentByte(file *os.File) []byte {
	var resultBuf []byte
	buffer := make([]byte, 100)
	for {
		readByte, err := file.Read(buffer)
		if err == io.EOF {
			resultBuf = append(resultBuf, buffer[:readByte]...)
			return resultBuf
		}
		if err != nil {
			return nil
		}
		resultBuf = append(resultBuf, buffer[:readByte]...)
	}
}

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	var resultSlice []string
	fileInput, err := os.OpenFile(inputFileName, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(fileInput)
	defer fileInput.Close()
	for fileScanner.Scan() {
		stringInFile := strings.Split(fileScanner.Text(), " ")
		dateInString, errTime := time.Parse("02.01.2006", stringInFile[0])
		if errTime != nil {
			return nil, errTime
		}
		if start.Before(dateInString) && end.After(dateInString)  || start.Equal(dateInString) || end.Equal(dateInString){
			resultSlice = append(resultSlice, fileScanner.Text())
		}
	}
	if len(resultSlice) == 0 {
		return nil, errors.New("Нет строк")
	}
	return resultSlice, nil
}

func ModifyFile(filename string, pos int, val string) {
	fileMod, _ := os.OpenFile(filename, os.O_RDWR, 0600)
	defer fileMod.Close()
	fileMod.Seek(int64(pos), 0)
	fileMod.Write([]byte(val))
}

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	fileInput, errInpytFile := os.OpenFile(inputFilename, os.O_RDONLY, 0600)
	if errInpytFile != nil {
		return errInpytFile
	}
	fileInput.Seek(int64(startpos), 0)
	defer fileInput.Close()
	fileOut, errOutFile := os.Create(outFileName)
	defer fileOut.Close()
	if errOutFile != nil {
		return errOutFile
	}
	buffer := readContentByte(fileInput)
	_, err := fileOut.Write(buffer)
	if err != nil {
		return err
	}
	return nil
}

func LineByNum(inputFilename string, lineNum int) string {
	file, err := os.Open(inputFilename)
	defer file.Close()
	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(file)
	stringNum := 0
	for fileScanner.Scan() {
		if stringNum == lineNum {
			return fileScanner.Text()
		}
		stringNum++
	}
	return ""
}

func main() {
	Println(ExtractLog("C:\\Users\\kss\\Desktop\\Новая папка\\LearnGo\\Go_second\\Files\\file1.txt", time.Date(2022, 12, 14, 0, 0, 0, 0, time.UTC), time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC)))
}
