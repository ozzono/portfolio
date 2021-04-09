package requests

import (
	"api-01/models"
	"api-01/utils"
	"log"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	testUser = models.User{
		Email:    "test@me.more",
		Password: "testing time",
	}
	startUser = models.User{
		Name:     "api-01 Challenge",
		Email:    "challenge@me.more",
		Password: "winner lottery ticket",
	}
	testReq   models.ActivationRequest
	testToken string
)

func init() {
	os.Setenv("MONGOHOSTNAME", "localhost")
	_, err := models.DefaultDB()
	if err != nil {
		log.Fatal(err)
	}
}

type custom struct {
	*testing.T
}

//Although very good, the idea of making such a kind of test never occurred to me before
func TestPing(t *testing.T) {
	err := Ping()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLogin(t *testing.T) {
	t.Logf("startUser.Email   : %s", startUser.Email)
	t.Logf("startUser.Password: %s", startUser.Password)
	outToken, err := Login(startUser)
	if err != nil {
		t.Fatal(err)
	}
	testToken = outToken
}

func TestAddUser(t *testing.T) {
	err := AddUser(testUser, testToken)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUser(t *testing.T) {
	outUser, err := GetUser(testUser, testToken)
	if err != nil {
		t.Fatal(err)
	}
	testUser = outUser
}

func TestUpdateUser(t *testing.T) {
	testUser.Password = utils.RString(6, 10)
	err := UpdateUser(testUser, testToken)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdd(t *testing.T) {
	reqID, err := AddReq(testUser.ID.Hex(), testToken)
	if err != nil {
		t.Fatal(err)
	}

	reqOID, err := primitive.ObjectIDFromHex(reqID)
	if err != nil {
		t.Fatal(err)
	}
	testReq.ID = reqOID
}

func TestGet(t *testing.T) {
	outReq, err := GetReq(testReq.ID.Hex(), testToken)
	if err != nil {
		t.Fatal(err)
	}
	testReq = outReq
}

func TestUpdateRequest(t *testing.T) {
	for key := range models.StatusEnum {
		err := UpdateReq(testReq.ID.Hex(), testUser.ID.Hex(), key, testToken)
		if err != nil {
			t.Fatal(err)
		}
	}
}
