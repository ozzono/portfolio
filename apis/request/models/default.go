package models

import (
	"fmt"
	db "request/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const api01DB = "request"

var DefaultUser = User{
	Name:     "request Challenge",
	Email:    "challenge@me.more",
	Password: "winner lottery ticket",
}

// DefaultDB ...
func DefaultDB() (db.Client, error) {
	client, err := db.NewClient()
	if err != nil {
		return db.Client{}, fmt.Errorf("db.NewClient err: %v", err)
	}

	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)

	if err := client.C.Database(api01DB).Drop(client.Ctx); err != nil {
		return db.Client{}, fmt.Errorf("api01DB.Drop err: %v", err)
	}

	if err := DefaultUser.Add(); err != nil {
		return db.Client{}, fmt.Errorf("Add err: %v", err)
	}

	req := ActivationRequest{
		ID:        primitive.NewObjectID(),
		Requestee: DefaultUser.ID,
		Approver:  DefaultUser.ID,
		Status:    "pending",
	}
	_, err = req.Add()
	if err != nil {
		return db.Client{}, fmt.Errorf("req.Add err: %v", err)
	}

	return client, nil
}
