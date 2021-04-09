package route

import (
	"api-01/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 2. solicita uma ativação de débito automático 	- add
// 3. cancela uma solicitação de ativação 			- update
// 4. aprova uma solicitação de ativação 			- update
// 5. rejeita uma solicitação de ativação 			- update
// 6. visualiza uma solicitação 					- get

func AddRequest(w http.ResponseWriter, r *http.Request) {
	userID, _, err := tellMeMore(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%s\"}", err), http.StatusBadRequest)
		return
	}

	if len(userID) == 0 {
		http.Error(w, fmt.Sprint("{\"msg\":\"user_id cannot be empty\"}"), http.StatusBadRequest)
		return
	}

	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	req := models.ActivationRequest{
		Requestee: userOID,
	}
	reqID, err := req.Add()
	if err != nil {
		log.Printf("req.Add err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}

	payload, err := json.Marshal(map[string]string{
		"success":    "successfully added request",
		"request_id": reqID,
	})
	if err != nil {
		log.Printf("json.Marshal err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	http.Error(w, string(payload), http.StatusOK)
}

func UpdateRequest(w http.ResponseWriter, r *http.Request) {
	userID, reqID, err := tellMeMore(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%s\"}", err), http.StatusBadRequest)
		return
	}

	if len(userID) == 0 {
		http.Error(w, fmt.Sprint("{\"msg\":\"user_id cannot be empty\"}"), http.StatusBadRequest)
		return
	}

	if len(reqID) == 0 {
		http.Error(w, fmt.Sprint("{\"msg\":\"request_id cannot be empty\"}"), http.StatusBadRequest)
		return
	}

	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}

	reqOID, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	url := strings.Split(strings.Split(r.RequestURI, "?")[0], "/")
	req := models.ActivationRequest{
		ID:       reqOID,
		Approver: userOID,
		Status:   models.StatusEnum[url[len(url)-1]],
	}
	err = req.Update(
		bson.D{
			{"$set", bson.D{{"status", req.Status}}},
			{"$set", bson.D{{"approver", req.Approver}}},
		},
	)
	if err != nil {
		log.Printf("req.Add err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	http.Error(w, "{\"success\":\"user successfully added\"}", http.StatusOK)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	_, reqID, err := tellMeMore(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"msg\":\"%s\"}", err), http.StatusBadRequest)
		return
	}
	reqOID, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		http.Error(w, "{\"msg\":\"invalid request id\"}", http.StatusBadRequest)
		return
	}
	req := models.ActivationRequest{
		ID: reqOID,
	}
	req, found, err := req.Find()
	if err != nil {
		log.Printf("req.Find err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	if !found {
		http.Error(w, "{\"msg\":\"request not found\"}", http.StatusOK)
		return
	}
	byteReq, err := json.Marshal(req)
	if err != nil {
		log.Printf("json.Marshal err: %v", err)
		http.Error(w, "{\"msg\":\"internal err; contact system admin\"}", http.StatusBadRequest)
		return
	}
	http.Error(w, string(byteReq), http.StatusOK)
}

func tellMeMore(r *http.Request) (string, string, error) {
	payload, err := PayTheLoad(r)
	if err != nil {
		return "", "", fmt.Errorf("invalid json data")
	}

	err = OpenTheGates(payload["token"])
	if err != nil {
		return "", "", fmt.Errorf("unauthorized access")
	}

	return payload["user_id"], payload["request_id"], nil
}
