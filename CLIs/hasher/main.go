package main

import (
	"hasher/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run - %v", err)
	}
}
