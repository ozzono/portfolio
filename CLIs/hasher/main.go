package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/melbahja/got"
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
	if err := config.DownloadFromURL(); err != nil {
		log.Fatalf("config.DownloadFromURL - %v", err)
	}

	config.SwiftDL()
}

func (c config) DownloadFromURL() error {
	log.Println("DownloadFromURL")
	t1 := time.Now()
	defer func() {
		log.Printf("DownloadFromURL duration %dms", time.Since(t1).Milliseconds())
	}()

	// create client
	client := c.NewClient()
	req, err := grab.NewRequest(tmpData, c.Url)
	if err != nil {
		return errors.Wrap(err, "grab.NewRequest")
	}
	defer cleanTmp()

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := transfer{client.Do(req)}
	fmt.Printf("- %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(time.Duration(c.SleepTicker) * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			resp.Log()
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		return errors.Wrapf(err, "Download failed: %v\n")
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	return nil
}

func (c config) SwiftDL() {
	//under test
	log.Println("DownloadFromURL")
	t1 := time.Now()
	defer func() {
		log.Printf("DownloadFromURL duration %dms", time.Since(t1).Milliseconds())
	}()

	g := got.New()

	err := g.Download(c.Url, tmpData)

	if err != nil {
		log.Println(err)
	}
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

func cleanTmp() {
	if err := os.Remove(tmpData); err != nil {
		log.Printf("os.Remove - %v", err)
	}
}

type transfer struct {
	*grab.Response
}

func (t transfer) Log() {
	charSize := fmt.Sprint(len(fmt.Sprint(t.Size())))
	format := "- transferred %0" + charSize + "v / %v bytes - % .2f%%\n"
	fmt.Printf(format, t.BytesComplete(), t.Size(), 100*t.Progress())
}
