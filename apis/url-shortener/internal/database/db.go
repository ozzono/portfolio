package database

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client has mongo client and context cancel function
type Client struct {
	C   *mongo.Client
	Ctx context.Context
}

const uri = "mongodb://%s:27017/?readPreference=primary&ssl=false"

// returns a mongodb client and pings the established connection
func NewClient() (*Client, error) {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf(uri, os.Getenv("MONGOHOSTNAME"))))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Connect")
	}
	// Ping the primary
	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, "client.C.Ping")
	}

	return &Client{C: c, Ctx: context.Background()}, nil
}
