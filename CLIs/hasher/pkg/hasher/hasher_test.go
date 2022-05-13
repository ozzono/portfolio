package hasher

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash1(t *testing.T) {
	hashed, err := Hasher(string([]byte{12}))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	expectedHash := []byte{0, 101, 142, 18, 48, 112, 197, 47}
	expectedHexa := "00658e123070c52f"

	assert.EqualValues(t, expectedHash, hashed, "returned hashed value different from expected")
	assert.EqualValues(t, expectedHexa, hex.EncodeToString(hashed), "returned hashed value different from expected")

	t.Log(hashed)
	t.Logf("encoded to hexa: %v", hex.EncodeToString(hashed))
}

func TestHash2(t *testing.T) {
	hashed, err := Hasher("mystical data to be hashed")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	expectedHash := []byte{0, 141, 192, 24, 34, 25, 192, 167}
	expectedHexa := "008dc0182219c0a7"

	assert.EqualValues(t, expectedHash, hashed, "returned hashed value different from expected")
	assert.EqualValues(t, expectedHexa, hex.EncodeToString(hashed), "returned hashed value different from expected")

	t.Log(hashed)
	t.Logf("encoded to hexa: %v", hex.EncodeToString(hashed))
}
