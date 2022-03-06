package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func ReadJson(path string, data interface{}) error {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "ioutil.ReadFile err")
	}

	if err = json.Unmarshal(jsonFile, &data); err != nil {
		return errors.Wrap(err, "json.Unmarshal err")
	}

	return nil
}
