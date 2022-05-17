package main

import "rest-get-api/pkg/handler"

func main() {
	handler := handler.Handler()
	handler.Handle("/", handler)
}
