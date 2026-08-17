package main

import (
	"context"
	"fmt"
	"io"
	"strings"
	"bytes"
)

func Contains(ctx context.Context,r io.Reader, seq []byte) (bool, error) {
    buff := make([]byte, len(seq))
    // Начальное чтение данных в буфер
    n, err := r.Read(buff)
    if err != nil && err != io.EOF {
        return false, err
    }
    if n != len(seq) {
        return false, nil
    }
    // Процесс поиска в потоке данных
    for {
        select {
		case <-ctx.Done():
			// Если контекст отменен, возвращаем ошибку отмены
			return false, ctx.Err()
		default:
        }
        if bytes.Equal(seq, buff) {
            return true, nil
        }
        buff = append(buff[1:], 0) // Сдвиг в буфере
        _, err := r.Read(buff[len(buff)-1:])
        if err != nil {
            return false, nil
        }
    }
}

func main() {
	ctx := context.Background()
	reader := strings.NewReader("43253532432432423")
	fmt.Println(Contains(ctx, reader, []byte("345")))
}
