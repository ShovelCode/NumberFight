package main

import (
	"fmt"
	"sync"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	var wg sync.WaitGroup

	numberF := 10 // starting value
	numberJ := 10 // starting value

	fChan := make(chan int)
	jChan := make(chan int)

	go func() {
		defer wg.Done()
		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if char == 'f' || char == 'F' {
				numberF--
				fChan <- numberF
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if char == 'j' || char == 'J' {
				numberJ--
				jChan <- numberJ
			}
		}
	}()

	wg.Add(2)

	go func() {
		for {
			select {
			case newF := <-fChan:
				fmt.Printf("Number F: %d\n", newF)
			case newJ := <-jChan:
				fmt.Printf("Number J: %d\n", newJ)
			}
		}
	}()

	wg.Wait()
}
