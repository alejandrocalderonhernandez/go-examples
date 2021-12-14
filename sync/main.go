package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 1000
)

//Only can be one process writing
func Deposit(ammount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock() //block variable access
	b := balance
	balance = ammount + b
	lock.Unlock() //After finish unlock variable access
}

//Can be one many processes reading
func Balance(lock *sync.RWMutex) int {
	lock.RLock() //Permit have many readers without blook line 24
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	var l sync.RWMutex //Objecto to lock and unlock to read and writh
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit((1000 * i), &wg, &l)
	}
	wg.Wait()
	fmt.Println(Balance(&l))
}
