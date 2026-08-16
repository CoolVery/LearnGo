package main

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (s *SafeMap) Get(key string) interface{} {
	var result interface{}
	s.mux.Lock()
	result = s.m[key]
	s.mux.Unlock()
	if result == nil {
		return nil
	}
	return result
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}

// -----------------------------------

type Counter struct {
    value int
    mu    sync.RWMutex
}

type Сount interface{
    Increment() // увеличение счётчика на единицу
    GetValue() int // получение текущего значения
}

func (counter *Counter) Increment() {
	counter.mu.Lock()
	counter.value++
	counter.mu.Unlock()
}

func (counter *Counter) GetValue() int {
	var result int
	counter.mu.RLock()
	result = counter.value
	counter.mu.RUnlock()
	return result
}

//-------------------------------------

type ConcurrentQueue struct {
    queue []interface{}
    mutex sync.Mutex
}

type Queue interface {
    Enqueue(element interface{}) // положить элемент в очередь
    Dequeue() interface{} // забрать первый элемент из очереди
}

func (cQueue *ConcurrentQueue) Enqueue(element interface{}) {
	cQueue.mutex.Lock()
	cQueue.queue = append(cQueue.queue, element)
	cQueue.mutex.Unlock()
}

func (cQueue *ConcurrentQueue) Dequeue() interface{} {
	var result interface{}
	cQueue.mutex.Lock()
	result = cQueue.queue[0]
	cQueue.queue = cQueue.queue[1:]
	cQueue.mutex.Unlock()
	if result == nil {
		return nil
	}
	return result
}

//-------------------------

var (
    Buf   []int
    mutex sync.Mutex
)

func Write(num int) {
	mutex.Lock()
	Buf = append(Buf, num)
	mutex.Unlock()
}

func Consume() int {
	var result int
	mutex.Lock()
	result = Buf[0]
	mutex.Unlock()
	return result
}