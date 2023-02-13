package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t1 := Now()
	time.Sleep(time.Second)
	t2, _ := time.Parse(TimeFormat, time.Now().Format(TimeFormat))
	assert.True(t, t1.Before(t2))
}
