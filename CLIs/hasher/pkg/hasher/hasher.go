package hasher

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"

	"github.com/pkg/errors"
)

var (
	coefficients = [8]int{2, 3, 5, 7, 11, 13, 17, 19}
)

type Hashed struct {
	Hash []byte
	Hex  string
}

func Hasher(input string) (*Hashed, error) {
	hashedInput := []byte{}
	for _, b := range sha256.Sum256([]byte(input)) {
		hashedInput = append(hashedInput, b)
	}

	reduced, err := reduceBytes(hashedInput)
	if err != nil {
		return nil, errors.Wrap(err, "reduceBytes")
	}

	custom := customHasher(reduced)
	return &Hashed{
		Hash: custom,
		Hex:  hex.EncodeToString(custom),
	}, nil
}

// I did not get what was meant to be done here
//
// -- where does the incoming byte come from?
// ---- which hash algorythm should be used?
func customHasher(input []byte) (output []byte) {
	// for each incoming byte, ib:
	// for each byte of the hash, h
	// h[i] = ((h[i-1] + ib) * coefficient[i]) % 255
	// in the case where i-1 == -1, h[i-1] should be 0.
	for i := range input {
		var h byte
		if i == 0 {
			h = 0
		} else {
			h = input[i-1]
		}
		output = append(output, (h*byte(coefficients[i]))%255)
	}
	return
}

func reduceBytes(input []byte) ([]byte, error) {
	numberInput := binary.LittleEndian.Uint64(input)
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, numberInput>>1); err != nil {
		return nil, errors.Wrap(err, "binary.Write")
	}
	if len(buf.Bytes()) > 8 {
		return reduceBytes(buf.Bytes())
	} else {
		return buf.Bytes(), nil
	}
}
