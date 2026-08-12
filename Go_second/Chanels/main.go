package main

import (
	"fmt"
	"reflect"
)

func SendFirst(ch chan int, num int) {
	ch <- num
}

func Receive(ch chan int) int {
	result := <- ch
	return result
}

func Send(ch1, ch2 chan int) {
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		ch2 <- 0
		ch2 <- 1
		ch2 <- 2
	}()
}

func Process(nums []int) chan int {
	resultCh := make(chan int, 10)
	for i := 0; i < len(nums); i++ {
		resultCh <- nums[i]
	}
	return resultCh
}

func main() {
	ch := make(chan int)
	go SendFirst(ch, 4)
	val := Receive(ch)
	fmt.Println(val)
	
}
