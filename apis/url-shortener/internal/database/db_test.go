package database

import (
	"context"
	"os"
	"testing"
)

func TestLocalHostDB(t *testing.T) {
	os.Setenv("MONGOHOSTNAME", "localhost")
	client, err := NewClient()
	if err != nil {
		t.Log("NewClient", err)
		t.FailNow()
	}
	// defer client.Cancel()
	defer func() {
		if err := client.C.Disconnect(context.Background()); err != nil {
			t.Log("client.Disconnect", err)
			t.Fail()
		}
	}()
	if err != nil {
		t.Fatalf("Datase err: %v", err)
	}
}
