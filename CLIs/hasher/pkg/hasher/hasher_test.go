package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash1(t *testing.T) {
	hashed, err := Hasher(string([]byte{12}))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	expected := &Hashed{
		Hash: []byte{0, 101, 142, 18, 48, 112, 197, 47},
		Hex:  "00658e123070c52f"}

	assert.EqualValues(t, expected.Hash, hashed.Hash, "returned hashed value different from expected")
	assert.EqualValues(t, expected.Hex, hashed.Hex, "returned hashed value different from expected")

	t.Logf("hashed.Hash -- %v", expected.Hash)
	t.Logf("hashed.Hex --- %v", expected.Hex)
}

func TestHash2(t *testing.T) {
	hashed, err := Hasher("mystical data to be hashed")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	expected := &Hashed{
		Hash: []byte{0, 251, 85, 247, 166, 162, 45, 30},
		Hex:  "00fb55f7a6a22d1e"}

	assert.EqualValues(t, expected.Hash, hashed.Hash, "returned hashed value different from expected")
	assert.EqualValues(t, expected.Hex, hashed.Hex, "returned hashed value different from expected")

	t.Logf("hashed.Hash -- %v", hashed.Hash)
	t.Logf("hashed.Hex --- %v", hashed.Hex)
}
