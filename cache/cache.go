package main

import (
	"fmt"
	"log"
	"time"
)

//interface {} -> is like a generic type T

//complex function
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

//Model to memory cells
type Memory struct {
	f     Function       //to call function
	cache map[int]Result //To cache
}

//@Param key to call Fibonacci func
type Function func(key int) (interface{}, error)

//struct to save result
type Result struct {
	value interface{}
	err   error
}

//Constructor for Memory
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]Result),
	}
}

//@Return a result in chahce if not exit call a Fibonacci func
func (m *Memory) GetOrCall(key int) (interface{}, error) {
	result, exist := m.cache[key] //verify if exist a valie un map

	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}

	return result.value, result.err

}

//Call a fibonacci function
func GetFib(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFib)
	fibo := []int{25, 8, 25, 7, 6}
	for _, n := range fibo {
		start := time.Now()
		v, e := cache.GetOrCall(n)
		if e != nil {
			log.Println(e)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), v)
	}
}
