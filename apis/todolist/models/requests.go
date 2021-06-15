package models

import (
	"context"
	"fmt"
	"log"
	"strings"
	db "todo/database"
	"todo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 2. solicita uma ativação de débito automático
// 3. cancela uma solicitação de ativação
// 4. aprova uma solicitação de ativação
// 5. rejeita uma solicitação de ativação
// 6. visualiza uma solicitação

const (
	reqsColl = "activation_requests"
)

var (
	StatusEnum = map[string]string{
		"approve":   "approved",
		"unapprove": "unapproved",
		"cancel":    "cancelled",
	}
)

type ActivationRequest struct {
	ID        primitive.ObjectID `bson:"_id"       json:"id"`
	Requestee primitive.ObjectID `bson:"requestee" json:"requestee"`
	Approver  primitive.ObjectID `bson:"approver"  json:"approver"`
	Status    string             `bson:"status"    json:"status"`
	log       bool
}

func (req *ActivationRequest) objectify(requestee, approver string) error {
	oRequestee, err := primitive.ObjectIDFromHex(requestee)
	if err != nil {
		return fmt.Errorf("primitive.ObjectIDFromHex err: %v", err)
	}
	req.Requestee = oRequestee

	oApprover, err := primitive.ObjectIDFromHex(approver)
	if err != nil {
		return fmt.Errorf("primitive.ObjectIDFromHex err: %v", err)
	}
	req.Approver = oApprover

	return nil
}

// Add ...
func (req *ActivationRequest) Add() (string, error) {
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return "", fmt.Errorf("db.NewClient err: %v", err)
	}

	req.ID = primitive.NewObjectID()
	req.Status = "pending"
	bsonReq, err := utils.ToDoc(req)
	if err != nil {
		return "", fmt.Errorf("req.toDoc err: %v", err)
	}

	req.Log("add")
	reqColl := client.C.Database(api01DB).Collection(reqsColl)
	_, err = reqColl.InsertOne(client.Ctx, bsonReq)
	if err != nil {
		return "", fmt.Errorf("reqColl.InsertOne err: %v", err)
	}
	return req.ID.Hex(), nil
}

func (req *ActivationRequest) Delete() error {
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return fmt.Errorf("db.NewClient err: %v", err)
	}

	reqsCollection := client.C.Database(api01DB).Collection(reqsColl)
	_, err = reqsCollection.DeleteOne(context.TODO(), bson.M{"_id": req.ID})
	if err != nil {
		return fmt.Errorf("reqsCollection.DeleteOne err: %v", err)
	}
	return nil
}

func (req *ActivationRequest) Find() (ActivationRequest, bool, error) {
	req.Log("searching")
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return ActivationRequest{}, false, fmt.Errorf("db.NewClient err: %v", err)
	}

	cursor, err := client.C.Database(api01DB).Collection(reqsColl).Find(client.Ctx, bson.M{"_id": req.ID})
	if err != nil {
		return ActivationRequest{}, false, fmt.Errorf("reqsCollection.Find err: %v", err)
	}
	reqs := []ActivationRequest{}
	for cursor.Next(client.Ctx) {
		r := ActivationRequest{}
		err = cursor.Decode(&r)
		if err != nil {
			return ActivationRequest{}, false, fmt.Errorf("cursor.Decode err: %v", err)
		}
		reqs = append(reqs, r)
	}

	if len(reqs) == 0 {
		return ActivationRequest{}, false, nil
	}

	return reqs[0], true, nil
}

func (req *ActivationRequest) Update(update bson.D) error {
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return fmt.Errorf("db.NewClient err: %v", err)
	}

	// for the sake of the challenge, I won't validate if the test was previously approved
	// depending of business rules in real applications, such validation would be needed or considered
	_, err = client.C.Database(api01DB).Collection(reqsColl).UpdateOne(
		client.Ctx,
		bson.M{"_id": req.ID},
		update,
	)
	if err != nil {
		return fmt.Errorf("reqsCollection.ReplaceOne err: %v", err)
	}
	return nil
}

func (req *ActivationRequest) Log(statement string) {
	if !req.log && !strings.Contains(statement, "debug") {
		return
	}
	log.Printf("%s - request\n", statement)
	log.Printf("id        : %v", req.ID.Hex())
	log.Printf("requestee : %v", req.Requestee.Hex())
	log.Printf("approver  : %v", req.Approver.Hex())
	log.Printf("status    : %v", req.Status)
}
