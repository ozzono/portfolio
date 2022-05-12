package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	dl "hasher/pkg/download"
	// "hasher/pkg/utils"

	"github.com/cavaliergopher/grab/v3"
	"github.com/pkg/errors"
)

const (
	tmpData         = ".data"
	highBuffSize    = 1024
	highSleepTicker = 1024
)

var (
	throttle      bool
	dest          string
	targetURL     string
	sleepTicker   int
	sizeLimit     int
	configPath    string
	errInvalidArg = fmt.Errorf("invalid argument")
)

type config struct {
	Url         string `json:"url"`
	SleepTicker int    `json:"sleepTicker"`
	SizeLimit   int    `json:"sizeLimit"`
	Throttle    bool   `json:"throttle"`
	DestPath    string `json:"destPath"`
	Debug       bool   `json:"debug"`
}

func init() {
	flag.StringVar(&configPath, "c", "./samples/config_sample-1.json", "config json file path - has priority over other arguments")
	flag.StringVar(&dest, "dest", "", "destiny file path")
	flag.StringVar(&targetURL, "url", "", "url value")
	flag.BoolVar(&throttle, "throttle", false, "enables download throttling")
	flag.IntVar(&sleepTicker, "sleep-ticker", 500, "sleep ticker")
	flag.IntVar(&sizeLimit, "size-limit", highBuffSize, "size limit")
}

func main() {
	if err := validateArgs(); err != nil {
		log.Fatalf("validateArgs - %v", err)
	}
	config, err := getConfig()
	if err != nil {
		log.Fatalf("getConfig - %v", err)
	}

	config.Log()
	dl := dl.DL{
		Dest:     config.DestPath,
		URL:      config.Url,
		Throttle: config.Throttle,
		Debug:    config.Debug,
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		t, err := dl.GotPkgDL()
		if err != nil {
			log.Printf("dl.GotPkgDL - %v", err)
		}
		log.Printf("dl.GotPkgDL duration - %dms", t)
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		t, err := dl.GrabPkgDL()
		if err != nil {
			log.Printf("dl.GrabPkgDL - %v", err)
		}
		log.Printf("dl.GrabPkgDL duration - %dms", t)
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		t, err := dl.StdLibDL()
		if err != nil {
			log.Printf("dl.StdLibDL - %v", err)
		}
		log.Printf("dl.StdLibDL duration - %dms", t)
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

func (c config) NewClient() *grab.Client {
	client := grab.NewClient()
	if c.Throttle {
		log.Printf("using hight buffer size of %d", c.SizeLimit)
		client.BufferSize = c.SizeLimit * 1024
	}
	return client
}

func getConfig() (*config, error) {

	if configPath == "" {
		return &config{
			Url:         targetURL,
			SleepTicker: sleepTicker,
			SizeLimit:   sizeLimit,
			Throttle:    throttle,
			DestPath:    dest,
		}, nil
	}

	byteData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadFile")
	}
	d := new(config)
	if err := json.Unmarshal(byteData, d); err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}
	return d, nil
}

func (c config) Log() {
	log.Printf("config.Url ---------- %v", c.Url)
	log.Printf("config.SleepTicker -- %v", c.SleepTicker)
	log.Printf("config.SizeLimit ---- %v", c.SizeLimit)
	log.Printf("config.Throttle ----- %v", c.Throttle)
}
