package db

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("MONGOHOSTNAME", "localhost")
}

func TestDatabase(t *testing.T) {
	client, err := NewClient()
	defer client.Cancel()
	defer client.C.Disconnect(client.Ctx)
	if err != nil {
		t.Fatalf("Datase err: %v", err)
	}
}
