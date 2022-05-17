package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"rest-get-api/pkg/database"

	"github.com/gorilla/mux"
)

func Handler() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ping", pong).Methods(http.MethodGet)
	r.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	return r
}

func pong(w http.ResponseWriter, r *http.Request) {
	log.Println("pong")
	w.Write([]byte("pong"))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if name != "" {
		log.Println("getting users")
	} else {
		log.Printf("getting user %s", name)
	}

	users := database.GetUsers(name)
	byteUsers, err := json.MarshalIndent(users, "", "	")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteUsers)
}
