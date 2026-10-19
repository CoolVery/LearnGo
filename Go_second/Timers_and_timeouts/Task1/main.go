package main

import (
	"errors"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	if n < 0 {
        return 0, errors.New("n must be non-negative")
    }
    if n < 2 {
        return n,nil
    }

	if timeout <= 0 {
		return 0, errors.New("timeout")
	}

	deadline := time.Now().Add(timeout)
	
    a, b := 0, 1
    for i := 2; i <= n; i++ {
		if time.Now().After(deadline) {
			return 0, errors.New("timeout")
		}
        a, b = b, a+b
    }

	
    return b, nil
}
