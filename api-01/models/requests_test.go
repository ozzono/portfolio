package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const reqDebug = false

func TestAddDelReq(t *testing.T) {
	req := ActivationRequest{
		Requestee: primitive.NewObjectID(),
		Approver:  primitive.NewObjectID(),
		log:       reqDebug,
	}
	reqID, err := req.Add()
	if err != nil {
		t.Fatalf("req.AddActivationRequest err: %v", err)
	}
	id, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		t.Fatalf("primitive.ObjectIDFromHex err: %v", err)
	}
	newReq := ActivationRequest{ID: id}
	_, found, err := newReq.Find()
	if err != nil {
		t.Fatalf("req.Find err: %v", err)
	}
	if !found {
		t.Fatal("request not found")
	}

	err = req.Delete()
	if err != nil {
		t.Fatalf("req.Delete err: %v", err)
	}

	_, found, err = req.Find()
	if err != nil {
		t.Fatalf("req.Find err: %v", err)
	}
	if found {
		t.Fatal("request found")
	}
}

func TestUpdateReq(t *testing.T) {
	req := ActivationRequest{
		ID:        primitive.NewObjectID(),
		Requestee: primitive.NewObjectID(),
		Status:    StatusEnum["approve"],
	}

	reqID, err := req.Add()
	if err != nil {
		t.Fatalf("req.AddActivationRequest err: %v", err)
	}
	id, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		t.Fatalf("primitive.ObjectIDFromHex err: %v", err)
	}
	newReq := ActivationRequest{ID: id, Approver: primitive.NewObjectID()}
	defer func() {
		err = newReq.Delete()
		if err != nil {
			t.Fatalf("newReq.Delete err: %v", err)
		}
	}()
	newReq, found, err := newReq.Find()
	if err != nil {
		t.Fatalf("failed to add request: %v", err)
	}
	if !found {
		t.Fatalf("failed to find new requests: %v", err)
	}

	err = newReq.Update(
		bson.D{
			{"$set", bson.D{{"status", StatusEnum["unapprove"]}}},
			{"$set", bson.D{{"approver", req.Approver}}},
		},
	)
	if err != nil {
		t.Fatalf("newReq.Update err: %v", err)
	}

	dbReq, found, err := newReq.Find()
	if err != nil {
		t.Fatalf("req.Find err: %v", err)
	}
	if !found {
		t.Fatal("request not found")
	}

	if dbReq.Status != StatusEnum["unapprove"] {
		t.Fatalf("failed to update; expected status: %s - returned status: %s", dbReq.Status, StatusEnum["unapprove"])
	}
	if dbReq.Approver.Hex() != req.Approver.Hex() {
		t.Fatalf("failed to update; expected approver: %s - returned approver: %s", dbReq.Approver.Hex(), newReq.Approver.Hex())
	}
}
