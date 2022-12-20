package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID        primitive.ObjectID `bson:"_id"       json:"id,omitempty"`
	Source    string             `bson:"source"    json:"source"`
	Shortened string             `bson:"shortened" json:"shortened,omitempty"`
	Count     int                `bson:"count"     json:"count,omitempty"`
}

func (url *URL) Log(header string, debug bool) {
	if !debug {
		return
	}
	if len(header) > 0 {
		log.Printf("%s - url\n", header)
		if url == nil {
			log.Println("nil url")
			return
		}
	}
	log.Printf("url.ID --------- %s", url.ID.Hex())
	log.Printf("url.Source ----- %s", url.Source)
	log.Printf("url.Shortened -- %s", url.Shortened)
	log.Printf("url.Count ------ %d", url.Count)
}

type ErrMsg struct {
	Msg string `json:"msg"`
	Err error  `json:"error,omitempty"`
}

func HTTPErr(c *gin.Context, msg ErrMsg, code int, err error) {
	logMsg := fmt.Sprintf("code: %d msg: %s", code, msg.Msg)
	if err != nil {
		logMsg += fmt.Sprintf(" err: %v", err)
	}
	log.Println(logMsg)

	c.JSON(
		code,
		msg,
	)
}

func URLToReader(source string) io.Reader {
	jsonURL, _ := json.Marshal(URL{Source: source})
	return bytes.NewReader(jsonURL)
}
