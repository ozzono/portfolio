package utils

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func RMFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return errors.Wrap(err, "os.Remove")
	}
	return nil
}

func ReadNEraseFile(filePath string) (string, error) {
	defer RMFile(filePath)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func WriteToFile(filePath string, data string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
