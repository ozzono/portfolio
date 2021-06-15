package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client has mongo client and context cancel function
type Client struct {
	C      *mongo.Client
	Cancel func()
	Ctx    context.Context
}

// Database returns a mongodb client and pings the established connection
func NewClient() (client Client, err error) {
	host := os.Getenv("MONGOHOSTNAME")
	host = "mongodb://root:challenge@" + host + ":27017/?authSource=admin&readPreference=primary&appname=mongodb-vscode%200.5.0&ssl=false"
	client.C, err = mongo.NewClient(
		options.Client().ApplyURI(host),
	)
	if err != nil {
		return
	}
	client.Ctx, client.Cancel = context.WithTimeout(context.Background(), 10*time.Second)

	err = client.C.Connect(client.Ctx)
	if err != nil {
		return
	}

	err = client.C.Ping(client.Ctx, readpref.Primary())
	if err != nil {
		return
	}

	return
}
