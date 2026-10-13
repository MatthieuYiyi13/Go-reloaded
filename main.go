package main

import (
	"fmt"
	"os"
	"reloaded/src"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: go run . <input> <output>")
	}
	input, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("read input: %w", err)
	}
	if err := os.WriteFile(args[1], []byte(reloaded.Up(string(input))), 0644); err != nil {
		return fmt.Errorf("write output: %w", err)
	}
	if err := os.WriteFile(args[1], []byte(reloaded.Low(string(input))), 0644); err != nil {
		return fmt.Errorf("write output: %w", err)
	}
	return nil
}
