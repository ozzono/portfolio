package main

import (
	"api-01/models"
	"api-01/route"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	mngUser        bool
	magicSecretKey = "I'm really enjoying this challenge"
	docker         bool
)

func init() {
	flag.BoolVar(&docker, "docker", false, "Redirects all requests using docker path to mongodb")
}

func main() {
	flag.Parse()
	if docker {
		os.Setenv("MONGOHOSTNAME", "mongodb")
	} else {
		os.Setenv("MONGOHOSTNAME", "localhost")
	}
	routes()
}

func routes() {
	_, err := models.DefaultDB()
	if err != nil {
		log.Fatalf("makeDefaultDB err: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	}).Methods("GET")

	// user routes
	r.HandleFunc("/user/add", route.AddUser).Methods("POST")
	r.HandleFunc("/user/login", route.LoginUser).Methods("POST")
	r.HandleFunc("/user/update", route.UpdateUser).Methods("POST")
	r.HandleFunc("/user/get", route.GetUser).Methods("POST")

	// request routes
	r.HandleFunc("/request/add", route.AddRequest).Methods("POST")
	r.HandleFunc("/request/get", route.GetRequest).Methods("POST")
	for key := range models.StatusEnum {
		r.HandleFunc("/request/"+key, route.UpdateRequest).Methods("POST")
	}

	r.HandleFunc("/refreshdb", route.RefreshDB).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
