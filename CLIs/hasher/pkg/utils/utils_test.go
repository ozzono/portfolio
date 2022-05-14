package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testFile = ".testfile"
	testData = "testData"
)

func TestFileMng(t *testing.T) {
	assert.NoError(t, WriteToFile(testFile, testData), "WriteToFile")
	_, err := ReadNEraseFile(testFile)
	assert.NoError(t, err, "ReadNEraseFile")
}
