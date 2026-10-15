package main

import (
	"fmt"
	"os"
	reloaded "reloaded/src"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}
	text := reloaded.Bin(string(input))
	text = reloaded.Hex(text)
	text = reloaded.Up(text)
	text = reloaded.Low(text)
	text = reloaded.Cap(text)
	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println("Erreur", err)
	}
}
