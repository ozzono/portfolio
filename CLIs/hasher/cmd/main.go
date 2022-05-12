package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	dl "hasher/pkg/download"

	"github.com/pkg/errors"
)

var (
	throttle      bool
	debug         bool
	dest          string
	targetURL     string
	configPath    string
	errInvalidArg = fmt.Errorf("invalid argument")
)

func init() {
	flag.StringVar(&configPath, "c", "./samples/config_sample-1.json", "config json file path - has priority over other arguments")
	flag.StringVar(&dest, "dest", "", "destiny file path")
	flag.StringVar(&targetURL, "url", "", "url value")
	flag.BoolVar(&throttle, "throttle", false, "enables download throttling")
	flag.BoolVar(&debug, "debug", false, "enables download debugging")
}

func Run() {
	if err := validateArgs(); err != nil {
		log.Fatalf("validateArgs - %v", err)
	}
	config, err := getConfig()
	if err != nil {
		log.Fatalf("getConfig - %v", err)
	}

	config.Log()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		t, err := config.GrabPkgDL()
		if err != nil {
			log.Printf("config.GrabPkgDL - %v", err)
		}
		log.Printf("config.GrabPkgDL duration - %dms", t)
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		t, err := config.GrabPkgDL()
		if err != nil {
			log.Printf("config.GrabPkgDL - %v", err)
		}
		log.Printf("config.GrabPkgDL duration - %dms", t)
		defer wg.Done()
	}()

	wg.Wait()
}

func validateArgs() error {
	flag.Parse()
	if configPath == "" {
		log.Println(`config file json not set; using the following arguments:
	-throttle
	-sleep-ticker
	-size-limit
	-dest
	-url
	if needed run with -h for usage details
		`)
		if dest == "" {
			return errors.Wrap(errInvalidArg, "`dest` cannot be empty")
		}
		if targetURL == "" {
			return errors.Wrap(errInvalidArg, "`url` cannot be empty")
		}
	}
	return nil
}

func getConfig() (*dl.Config, error) {

	if configPath == "" {
		return &dl.Config{
			URL:      targetURL,
			Throttle: throttle,
			Dest:     dest,
			Debug:    debug,
		}, nil
	}

	byteData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadFile")
	}
	d := new(dl.Config)
	if err := json.Unmarshal(byteData, d); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}
	return d, nil
}
