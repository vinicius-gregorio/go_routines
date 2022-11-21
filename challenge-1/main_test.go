package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello, world!"
	printMessage()
	_ = w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, world") {
		t.Errorf("Expected to find 'First print' in the output, but got %s", output)
	}

}

func Test_updateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("Hello, world!")

	wg.Wait()
	if msg != "Hello, world!" {
		t.Errorf("Expected msg to be 'Hello, world!', but got %s", msg)
	}
}
