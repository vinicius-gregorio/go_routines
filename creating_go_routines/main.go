package main

import (
	"fmt"
	"strconv"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
func main() {
	var wg sync.WaitGroup
	words := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	wg.Add(len(words))

	for i, v := range words {
		go printSomething(strconv.Itoa(i)+"  "+string(v), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("First print", &wg)
}
