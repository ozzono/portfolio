package models

import (
	"context"
	"fmt"
	"log"
	"reflect"

	db "api-01/database"
	"api-01/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	usersColl = "users"
)

// User defines the user in db
type User struct {
	ID          primitive.ObjectID `bson:"_id"   json:"id"`
	Name        string             `bson:"name"  json:"name"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password"`
	pwEncrypted bool
	log         bool
}

// Add creates a user record in the database
func (user *User) Add() error {
	user.Log("creating")
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return err
	}

	if len(user.Password) < 6 {
		return fmt.Errorf("password cannot be empty")
	}

	if !user.pwEncrypted {
		err = user.HashPassword(user.Password)
		if err != nil {
			return err
		}
	}

	dbUser, found, err := user.Find()
	if err != nil {
		return err
	}
	if found {
		if !reflect.DeepEqual(&dbUser, user) {
			err = user.HashPassword(user.Password)
			if err != nil {
				return err
			}

			err = user.Update()
			if err != nil {
				return err
			}

			return nil
		}
	}
	user.ID = primitive.NewObjectID()
	bsonUser, err := utils.ToDoc(user)
	if err != nil {
		return fmt.Errorf("utils.ToDoc err: %v", err)
	}
	userCollection := client.C.Database(api01DB).Collection(usersColl)
	_, err = userCollection.InsertOne(client.Ctx, bsonUser)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (user *User) Update() error {
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return fmt.Errorf("db.NewClient err: %v", err)
	}

	if !user.pwEncrypted {
		err = user.HashPassword(user.Password)
		if err != nil {
			return fmt.Errorf("HashPassword err: %v", err)
		}
	}
	_, err = client.C.Database(api01DB).Collection(usersColl).UpdateOne(
		client.Ctx,
		bson.M{"email": user.Email},
		bson.D{
			{"$set", bson.D{{"password", user.Password}}},
		},
	)
	if err != nil {
		return fmt.Errorf("usersCollection.ReplaceOne err: %v", err)
	}
	return nil
}

// Find searches the users collection using the email as key
func (user *User) Find() (User, bool, error) {
	user.Log("searching")
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return User{}, false, fmt.Errorf("db.NewClient err: %v", err)
	}

	usersCollection := client.C.Database(api01DB).Collection(usersColl)
	cursor, err := usersCollection.Find(client.Ctx, bson.M{"email": user.Email})
	if err != nil {
		return User{}, false, fmt.Errorf("usersCollection.Find err: %v", err)
	}
	users := []User{}
	for cursor.Next(client.Ctx) {
		u := User{}
		err = cursor.Decode(&u)
		if err != nil {
			return User{}, false, fmt.Errorf("cursor.Decode err: %v", err)
		}
		u.pwEncrypted = true
		users = append(users, u)
	}

	if len(users) == 0 {
		return User{}, false, nil
	}

	return users[0], true, nil
}

// Delete ...
func (user *User) Delete() error {
	user.Log("deleting")
	client, err := db.NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		return fmt.Errorf("db.NewClient err: %v", err)
	}

	usersCollection := client.C.Database(api01DB).Collection(usersColl)
	_, err = usersCollection.DeleteOne(context.TODO(), bson.M{"email": user.Email})
	if err != nil {
		return fmt.Errorf("usersCollection.DeleteOne err: %v", err)
	}
	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	user.pwEncrypted = true

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}

// Log ...
func (user *User) Log(header string) {
	if !user.log {
		return
	}
	if len(header) > 0 {
		fmt.Printf("%s - user\n", header)
	}
	log.Printf("ID        %s", user.ID.String())
	log.Printf("Name      %s", user.Name)
	log.Printf("Email     %s", user.Email)
	log.Printf("Password  %s", user.Password)
	log.Printf("Encrypted %t", user.pwEncrypted)
}
