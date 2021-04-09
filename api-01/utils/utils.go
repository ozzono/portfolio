package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	charSet = "abcdefghijklmnopqrstuwxyz"
)

func ToDoc(input interface{}) (bson.D, error) {
	data, err := bson.Marshal(input)
	if err != nil {
		return bson.D{}, fmt.Errorf("bson.Marshal err: %v", err)
	}

	doc := bson.D{}
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		return bson.D{}, fmt.Errorf("bson.Unmarshal err: %v", err)
	}
	return doc, nil
}

func RString(min, max int) string {
	out := ""
	for i := 0; i < (RInt(max) + min); i++ {
		out += string(charSet[RInt(len(charSet))])
	}
	return out
}

func RInt(i int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
}

func DefaultRequest(method, url, payload string) (*http.Response, error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(payload))
	// req, err := http.NewRequest()
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest err: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do err: %v", err)
	}
	return res, nil
}
