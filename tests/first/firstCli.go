package main

import (
	"fmt"

	"github.com/MikhailBatsin-code/librlcli/interactive"
)

func main() {
	interactive.AddAction("helloworld", "print hello world", func(args []string) error {
		fmt.Println("Hello world")

		return nil
	})

	interactive.AddAction("hello", "print hello to something or somebody", func(args []string) error {
		if len(args) == 1 {
			return fmt.Errorf("not enough arguments")
		}

		fmt.Println("Hello", args[1:])

		return nil
	})

	interactive.InteractiveCli()
}
