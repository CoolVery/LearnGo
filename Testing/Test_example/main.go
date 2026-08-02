package main

import "fmt"

func PrintHello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func Sum(a, b int) int {
	return a + b
}

func Length(a int) string {
    switch {
    case a < 0:
        return "negative"
    case a == 0:
        return "zero"
    case a < 10:
        return "short"
    case a < 100:
        return "long"
    }
    return "very long"
}

func Multiply(a, b int) int {
	return a + b
}