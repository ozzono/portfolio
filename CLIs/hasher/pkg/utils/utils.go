package utils

import (
	"log"
	"os"
)

func RMFile(path string) {
	if err := os.Remove(path); err != nil {
		log.Printf("os.Remove - %v", err)
	}
}
