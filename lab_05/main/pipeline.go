package main

import (
	"fmt"
	"sync"
)

func pipelineMtr(sliceMtr []matrixStruct, count, flag int) {
	inverseChan := make(chan matrixStruct, count)
	standartChan := make(chan matrixStruct, count)
	vinogradChan := make(chan matrixStruct, count)
	resultChan := make(chan matrixStruct, count)

	var wg sync.WaitGroup

	inverse := func() {
		defer wg.Done()
		for {
			select {
			case item, ok := <-inverseChan:
				if ok {
					item.inverseMatrix = gaussMethodEx(item.matrixA)
					standartChan <- item
				} else {
					close(standartChan)
					return
				}
			}
		}
	}

	standart := func() {
		defer wg.Done()
		for {
			select {
			case item, ok := <-standartChan:
				if ok {
					item.standartMul = mulStandart(item.inverseMatrix, item.matrixB, item.size)
					vinogradChan <- item
				} else {
					close(vinogradChan)
					return
				}
			}
		}
	}

	vinograd := func() {
		defer wg.Done()
		for {
			select {
			case item, ok := <-vinogradChan:
				if ok {
					item.vinogradMul = mulVinogradov(item.inverseMatrix, item.matrixB, item.size)
					resultChan <- item
				} else {
					close(resultChan)
					return
				}
			}
		}
	}

	go inverse()
	go standart()
	go vinograd()

	wg.Add(3)

	for i := 0; i < count; i++ {
		inverseChan <- sliceMtr[i]
	}

	close(inverseChan)
	wg.Wait()

	for item := range resultChan {
		fmt.Println(item)
		sliceMtr[item.uniq-1] = item
	}
}
