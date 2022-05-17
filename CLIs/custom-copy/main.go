package main

import (
	"flag"
	"fmt"
	"log"

	"custom-copy/pkg/utils"

	"github.com/pkg/errors"
)

var (
	source  string
	destiny string
	verbose bool
)

func init() {
	flag.StringVar(&source, "s", "", "Source directory to copy from")
	flag.StringVar(&destiny, "d", "", "Destiny directory to copy to")
	flag.BoolVar(&verbose, "v", false, "Enables verbose mode")
}

func main() {
	validateArgs()
	files, err := utils.ListPath(source, verbose)
	if err != nil {
		log.Println(err)
		return
	}

	if err := files.MkdirAll(source, destiny); err != nil {
		log.Panic(errors.Wrap(err, "files.MkdirAll"))
	}

	if err := files.CpAll(source, destiny); err != nil {
		log.Panic(errors.Wrap(err, "files.CpAll"))
	}
}

func validateArgs() error {
	flag.Parse()
	if destiny == "" {
		return fmt.Errorf("invalid destiny path; cannot be empty")
	}
	if source == "" {
		return fmt.Errorf("invalid source path; cannot be empty")
	}

	if verbose {
		fmt.Println("- verbose mode enabled")
	}
	return nil
}
