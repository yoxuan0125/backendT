package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var rng = rand.New(source)

func main() {
	var meats []Meat

	//進貨：牛肉10個，豬肉7個，雞肉5個
	for i := 0; i < 10; i++ {
		meats = append(meats, Meat{Type: "Beef", ProcessingTime: 1 * time.Second})
	}
	for i := 0; i < 7; i++ {
		meats = append(meats, Meat{Type: "Pork", ProcessingTime: 2 * time.Second})
	}
	for i := 0; i < 5; i++ {
		meats = append(meats, Meat{Type: "Chicken", ProcessingTime: 3 * time.Second})
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	// 創建五個員工
	employees := []Employee{
		{ID: "A"},
		{ID: "B"},
		{ID: "C"},
		{ID: "D"},
		{ID: "E"},
	}

	// 員工開始處理肉品
	for _, employee := range employees {
		wg.Add(1)
		go employee.HandlingMeat(&meats, &mu, &wg)
	}

	// 等待所有肉品處理
	wg.Wait()
	fmt.Println("全部肉品皆處理完畢")
}
