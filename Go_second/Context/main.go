package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
	"time"
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

//-------------------------------------

// type APIResponse struct {
//      Data string 
//      StatusCode int
// }

// func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error){
//     var resultDate APIResponse
//     ctxAPI, cancel := context.WithTimeout(ctx, timeout)
//     defer cancel()

//     client := http.Client{}

//     req, err := http.NewRequest("GET", url, nil)
//     if err != nil {
//         return nil, err
//     }
//     resp, err := client.Do(req.WithContext(ctxAPI))
//     if err != nil {
//         return nil, err
//     }

//     bodyBites, err := io.ReadAll(resp.Body)
//     if err != nil {
//         return nil, err
//     }
//     resultDate.Data = string(bodyBites)
//     resultDate.StatusCode = resp.StatusCode
//     defer resp.Body.Close()
//     return &resultDate, nil
// }

//----------------------------------------

type APIResponse struct { 
    URL string // запрошенный URL 
    Data string // тело ответа 
    StatusCode int // код ответа 
    Err error // ошибка, если возникла 
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
    ctxAPI, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    var wg sync.WaitGroup

    resultAPIResponse := make([]*APIResponse, 0)
    ch := make(chan *APIResponse, len(urls))
    client := &http.Client{}
    
    go func() {
        wg.Wait()
        close(ch)
    }()

    for _, url := range(urls) {
        select {
        case<-ctx.Done():
            return nil
        default:
            wg.Add(1)
            go getAPI(ctxAPI, url, client, ch, &wg)
        }
    }
    
    for resp := range ch {
        resultAPIResponse = append(resultAPIResponse, resp)
    }
    return resultAPIResponse
}

func getAPI(ctxAPI context.Context, url string, client *http.Client, ch chan *APIResponse, wg *sync.WaitGroup) {
    defer wg.Done()
    resultAPI := APIResponse {
        URL: "",
        Data: "",
        StatusCode: 0,
        Err: nil,
    }
 
    req, err := http.NewRequestWithContext(ctxAPI, "GET", url, nil)
    if err != nil {
        resultAPI.Err = err
        ch<-&resultAPI
        return
    }
    resp, err := client.Do(req.WithContext(ctxAPI))
    if err != nil {
        resultAPI.Err = err
        ch<-&resultAPI
        return
    }
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        resultAPI.Err = err
        ch<-&resultAPI
        return
    }
    resultAPI = APIResponse{
        URL: url,
        Data: string(bodyBytes),
        StatusCode: int(resp.StatusCode),
    }
    ch<-&resultAPI
    return
}

func StartServer(maxTimeout time.Duration) {
    
    mux := http.NewServeMux()

    mux.Handle("/readSource", getFromServer(maxTimeout))
    server := &http.Server{
		Addr:         ":8080",
		Handler:      mux, 
		ReadTimeout:  maxTimeout * 2,
		WriteTimeout: maxTimeout * 2,
	}

    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func getFromServer(maxTimeout time.Duration) http.Handler {
    innerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        innerReq, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/provideData", nil)
        if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
        client := &http.Client{}
        resp, err := client.Do(innerReq)
        if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
        w.WriteHeader(resp.StatusCode)
        io.Copy(w, resp.Body)
    })
    timeoutHandler := http.TimeoutHandler(innerHandler, maxTimeout, "")
    return timeoutHandler
}

func main() {
	// ctx := context.Background()
	// reader := strings.NewReader("43253532432432423")
	// fmt.Println(Contains(ctx, reader, []byte("345")))
    StartServer(5)
}
