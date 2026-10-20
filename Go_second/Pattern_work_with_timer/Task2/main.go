package main

import (
	"context"
	"io"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)
	select {
	case <- ctx.Done():
		return
	default:
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	buf, _ := io.ReadAll(f)
	
	select {
	case <- ctx.Done():
		return
	case result <- buf:
	}
}
