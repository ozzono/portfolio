package models

import (
	"os"
	"reflect"
	"testing"
)

const (
	userDebug = false
	bigSecret = "insert lottery number"
)

var testUser = User{
	Name:     "Test Me",
	Email:    "test@me.now",
	Password: bigSecret,
	log:      userDebug,
}

func init() {
	os.Setenv("MONGOHOSTNAME", "localhost")
}

func TestAddDelUser(t *testing.T) {
	err := testUser.Add()
	if err != nil {
		t.Fatalf("Add err: %v", err)
	}

	_, found, err := testUser.Find()
	if err != nil {
		t.Fatalf("FindUser err: %v", err)
	}
	if !found {
		t.Fatal("user not found")
	}

	err = testUser.Delete()
	if err != nil {
		t.Fatalf("Delete err: %v", err)
	}

	_, found, err = testUser.Find()
	if err != nil {
		t.Fatalf("FindUser err: %v", err)
	}
	if found {
		t.Fatal("user found")
	}

}

func TestUpdateUser(t *testing.T) {
	err := testUser.Add()
	if err != nil {
		t.Fatalf("Add err: %v", err)
	}

	_, found, err := testUser.Find()
	if err != nil {
		t.Fatalf("FindUser err: %v", err)
	}
	if !found {
		t.Fatal("user not found")
	}

	newSecret := "There's only one life-form as intelligent as me within thirty parsecs of here and that's me."
	// Marvin - Hitchhiker's Guide to the Galaxy

	err = testUser.HashPassword(newSecret)
	if err != nil {
		t.Fatalf("HashPassword err: %v", err)
	}

	err = testUser.Update()
	if err != nil {
		t.Fatalf("Update err: %v", err)
	}

	outUser, found, err := testUser.Find()
	if err != nil {
		t.Fatalf("Find err: %v", err)
	}
	if !found {
		t.Fatal("user not found")
	}

	err = outUser.CheckPassword(newSecret)
	if err != nil {
		t.Fatalf("Update password failure err: %v", err)
	}

	if !reflect.DeepEqual(
		User{Name: testUser.Name, Email: testUser.Email}, //before update
		User{Name: outUser.Name, Email: outUser.Email},   //after update
	) {
		t.Logf("expecting - name: %s email: %s", testUser.Name, testUser.Email)
		t.Logf("returned  - name: %s email: %s", outUser.Name, outUser.Email)
		t.Fatalf("Update password failure")
	}

	err = testUser.Delete()
	if err != nil {
		t.Fatalf("Delete err: %v", err)
	}

}

func TestCheckPassword(t *testing.T) {
	user := User{Password: bigSecret, log: true}
	err := testUser.HashPassword(user.Password)
	if err != nil {
		t.Fatalf("HashPassword err: %v", err)
	}
	err = testUser.CheckPassword(bigSecret)
	if err != nil {
		t.Fatalf("CheckPassword err: %v", err)
	}
}

func TestUserHashPassword(t *testing.T) {
	user := User{Password: bigSecret, log: true}
	err := testUser.HashPassword(user.Password)
	if err != nil {
		t.Fatalf("HashPassword err: %v", err)
	}
}
