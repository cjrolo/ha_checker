package main

import (
	"fmt"
)

func main() {
	backend, _ := NewBaseBackend("54.93.189.128", 8080)
	backend.Verify()
	fmt.Println("Is Jenkins down? ", backend.IsDown())
}
