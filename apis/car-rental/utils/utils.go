package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const TimeFormat = "2006-01-02 15:04:05"

type APIError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func (e APIError) LogTxt() string {
	return fmt.Sprintf(`
	ErrorCode ----- %d
	ErrorMessage -- %s
	`, e.ErrorCode, e.ErrorMessage)
}

func HTTPErrJSON(c *gin.Context, code int, msg string) {
	c.JSON(code, APIError{ErrorCode: code, ErrorMessage: msg})
}

// returns in `2023-02-10 01:47:01 +0000 UTC` format
//
// instead of `2023-02-10 01:47:01.8261636 +0000 UTC`
//
// if you say it is horrible, I'll agree
func Now() time.Time {
	t, _ := time.Parse(TimeFormat, time.Now().Format(TimeFormat))
	return t
}
