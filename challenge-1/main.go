package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

func printMessage(wg *sync.WaitGroup) {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	msg = "Hello, world!"

	phrases := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}

	wg.Add(len(phrases))

	for _, v := range phrases {
		go updateMessage(v)
		printMessage(&wg)

	}
	wg.Wait()

}
