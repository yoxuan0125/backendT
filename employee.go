package main

import (
	"fmt"
	"sync"
	"time"
)

type Employee struct {
	ID string
}

func (e Employee) HandlingMeat(meatChannel chan Meat, wg *sync.WaitGroup) {
	defer wg.Done()
	for meat := range meatChannel {
		fmt.Printf("%s 在 %s 取得 %s\n", e.ID, time.Now().Format("2006-01-02 15:04:05"), meat.Type)
		time.Sleep(meat.ProcessingTime) //模擬處理肉品
		fmt.Printf("%s 在 %s 處理完 %s\n", e.ID, time.Now().Format("2006-01-01 15:04:05"), meat.Type)
	}
}
