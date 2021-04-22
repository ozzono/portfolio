package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"request/auth"
	"request/models"

	"github.com/miguelpragier/handy"
)

const (
	minSize = 8
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user, err := comeJoinUs(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%s\"}", err), http.StatusBadRequest)
		return
	}

	err = user.Add()
	if err != nil {
		log.Printf("user.Add err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	http.Error(w, "{\"success\":\"user successfully added\"}", http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := comeJoinUs(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%s\"}", err), http.StatusBadRequest)
		return
	}

	user, found, err := user.Find()
	if err != nil {
		log.Printf("user.Add err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}

	if !found {
		http.Error(w, "{\"msg\":\"user not found\"}", http.StatusOK)
		return
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		log.Printf("json.Marshal err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	http.Error(w, string(byteUser), http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := comeJoinUs(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%s\"}", err), http.StatusBadRequest)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, fmt.Sprintf("{\"msg\":\"invalid password: %s; must have 6 or more characters\"}", err), http.StatusBadRequest)
		return
	}
	err = user.Update()
	if err != nil {
		log.Printf("user.Add err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	http.Error(w, "{\"success\":\"user successfully added\"}", http.StatusOK)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("json.NewDecoder err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	// I've added this verification trying to avoid NoSQL injection
	// Which my shallow testing and research showed unlikely to happen with Golang & MongoDB
	if !handy.CheckEmail(user.Email) {
		if err != nil {
			http.Error(w, "{\"msg\":\"invalid email format\"}", http.StatusBadRequest)
			return
		}
	}

	dbUser, found, err := user.Find()
	if err != nil {
		log.Printf("user.Find err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}

	if !found {
		http.Error(w, "{\"msg\":\"invalid email or password - 0\"}", http.StatusBadRequest)
		return
	}

	if err = dbUser.CheckPassword(user.Password); err != nil {
		http.Error(w, "{\"msg\":\"invalid email or password - 1\"}", http.StatusBadRequest)
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       magicSecretKey,
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	token, err := jwtWrapper.NewToken(user.Email)
	if err != nil {
		log.Printf("jwtWrapper.NewToken err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	output := map[string]string{
		"token": token,
	}
	byteLoad, err := json.Marshal(output)
	if err != nil {
		log.Printf("json.Marshal err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
	}
	http.Error(w, string(byteLoad), http.StatusOK)
}

func comeJoinUs(r *http.Request) (models.User, error) {
	payload, err := PayTheLoad(r)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid json data")
	}

	err = OpenTheGates(payload["token"])
	if err != nil {
		return models.User{}, fmt.Errorf("unauthorized access")
	}

	email, ok := payload["email"]
	if !ok {
		return models.User{}, fmt.Errorf("email cannot be empty")
	}

	return models.User{
		Email:    email,
		Password: payload["password"],
	}, nil
}
