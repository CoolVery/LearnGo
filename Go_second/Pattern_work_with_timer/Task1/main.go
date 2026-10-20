package main

import (
	"time"
)

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	defer close(prime_nums)

	timer := time.AfterFunc(time.Second / 10, func() {
		select {
		case stop <- struct{}{}:
		default:
		}
	})
	defer timer.Stop()

	for i := 2; i <= N; i++ {
		select {
		case <- stop:
			return
		default:
		}
		if !isPrime(i) {
			continue
		}
		select {
		case <- stop:
			return
		case prime_nums <- i:
		}
	}
}
