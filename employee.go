package main

import (
	"fmt"
	"sync"
	"time"
)

type Employee struct {
	ID string
}

func (e Employee) HandlingMeat(meatSlice *[]Meat, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mu.Lock() //利用互斥鎖確保每位員工取出肉品時，其他員工無法取出肉品
		if len(*meatSlice) == 0 {
			mu.Unlock()
			return
		}

		index := rng.Intn(len(*meatSlice)) //隨機選取肉品，並將其從肉品切片刪除
		meat := (*meatSlice)[index]
		*meatSlice = append((*meatSlice)[:index], (*meatSlice)[index+1:]...)
		mu.Unlock() //釋放鎖，讓其他人可以取得肉品

		fmt.Printf("%s 在 %s 取得 %s\n", e.ID, time.Now().Format("2006-01-02 15:04:05"), meat.Type)
		time.Sleep(meat.ProcessingTime)
		fmt.Printf("%s 在 %s 處理完 %s\n", e.ID, time.Now().Format("2006-01-01 15:04:05"), meat.Type)
	}
}
