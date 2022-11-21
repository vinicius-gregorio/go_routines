package main

import (
	"testing"
)

func Test_printMessage(t *testing.T) {
	// stdOut := os.Stdout

	// r, w, _ := os.Pipe()
	// os.Stdout = w

	// var wg sync.WaitGroup

	// wg.Add(1)

	// wg.Wait()

	// _ = w.Close()

	// res, _ := io.ReadAll(r)
	// output := string(res)

	// os.Stdout = stdOut

	// if !strings.Contains(output, "First print") {
	// 	t.Errorf("Expected to find 'First print' in the output, but got %s", output)
	// }

}

func Test_updateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("Hello, world!")

	wg.Wait()
	if msg != "Hello, world!" {
		t.Errorf("Expected msg to be 'Hello, world!', but got %s", msg)
	}
}
