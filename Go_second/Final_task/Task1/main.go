package main

import (
	"context"
	"io"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	newCtx, cancle := context.WithTimeout(ctx, timeout)
	defer cancle()

	resChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	go func(r io.Reader, res chan<- []byte, errCh chan<- error) {

		data, err := io.ReadAll(r)
		if err != nil {
			errChan <- err
			return
		}
		resChan <- data
	}(r, resChan, errChan)
	select {
	case <-newCtx.Done():
		return newCtx.Err()
	case err := <-errChan:
		return err
	// case data := <-resChan:
	// 	//TODO
	// }
	default:
		return nil
	}
}
