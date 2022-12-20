package main

import (
	"flag"
	"log"
	"os"
	"url-shortener/internal/handler"

	"github.com/pkg/errors"
)

var dockerized bool

func init() {
	flag.BoolVar(&dockerized, "docker", false, "indicates if code is being hosted locally of dockerized")
}

func main() {
	flag.Parse()
	if dockerized {
		os.Setenv("MONGOHOSTNAME", "mongodb")
	} else {
		os.Setenv("MONGOHOSTNAME", "localhost")
	}
	handler, err := handler.NewHandler()
	if err != nil {
		log.Fatal(errors.Wrap(err, "handler.NewHandler"))
	}
	handler.Router.Run(":8000")
}
