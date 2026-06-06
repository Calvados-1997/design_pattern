package main

import (
	"fmt"
	"os"
	"sample/builder"
)

func main() {
	pc, err := builder.NewPcBuilder().
		SetCPU("Intel").
		SetMemory(24).
		SetStorage(24).
		SetWifi(true).
		SetOS("Mac").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", fmt.Sprintln(err))
		return
	}
	fmt.Printf("created pc details: %+v\n", pc)
}
