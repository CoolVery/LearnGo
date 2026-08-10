package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	info  LogLevel = "Info"
	error LogLevel = "Error"
)

type Log struct {
	Level LogLevel
}

func (log Log) Log(message string) {
	switch log.Level {
	case info:
		fmt.Printf("INFO: %s.", message)
	case error:
		fmt.Printf("ERROR: %s.", message)
	}
}

func main() {
	// var c Shape = Circle{radius: 3}
	// var r Shape = Rectangle{width: 3, height: 2}
	// c.Area()
	// r.Area()

}
