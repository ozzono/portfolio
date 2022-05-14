package dl

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"hasher/pkg/utils"

	"github.com/cavaliergopher/grab/v3"
	"github.com/melbahja/got"
	"github.com/pkg/errors"
)

const (
	tmpData = ".data"
)

type Config struct {
	Dest        string `json:"dest"`
	URL         string `json:"url"`
	Throttle    bool   `json:"throttle"`
	Debug       bool   `json:"debug"`
	tmpFile     string
	ElapsedTime int64 //in millisecond
}

func (c *Config) TmpFile(path string) {
	c.tmpFile = path
}

type limit struct {
	ref string
}

func (c *Config) GrabPkgDL() (string, error) {
	log.Println("GrabDL")
	t1 := time.Now()
	defer func() {
		c.ElapsedTime = time.Since(t1).Milliseconds()
	}()

	suffix := "grab"
	if c.tmpFile != "" {
		suffix = c.tmpFile
	}

	client := grab.NewClient()

	req, err := grab.NewRequest(tmpData+suffix, c.URL)
	if err != nil {
		return "", errors.Wrap(err, "grab.NewRequest")
	}
	if c.Throttle {
		req.RateLimiter = limit{ref: suffix}
		req.RateLimiter.WaitN(req.Context(), 100)
	}

	if c.Debug {
		log.Printf("%s GrabDL Downloading %v", suffix, req.URL())
	}

	resp := transfer{client.Do(req), " " + suffix}
	if c.Debug {
		log.Printf("- %v", resp.HTTPResponse.Status)
	}

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			resp.Log(c.Debug)
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		return "", errors.Wrap(err, "GrabDL Download failed")
	}

	rawData, err := utils.ReadNEraseFile(tmpData + suffix)
	if err != nil {
		return "", errors.Wrap(err, "utils.ReadNEraseFile")
	}

	return rawData, nil
}

// This is a research leftover kept only because why not?
//
// note: this method was ditched as I didn't find throttling support
func (c *Config) GotPkgDL() (int64, error) {
	log.Println("GotPkgDL")
	t1 := time.Now()

	if err := got.New().Download(c.URL, tmpData+"gotpkgdl"); err != nil {
		return 0, errors.Wrap(err, "got.New().Download()")
	}

	return time.Since(t1).Milliseconds(), nil
}

// This is a research leftover kept only because why not?
//
// note: this method was ditched as I didn't find throttling support
func (c *Config) StdLibDL() (int64, error) {
	log.Println("StdLibDL")

	t1 := time.Now()
	out, err := os.Create(tmpData + "stdlibdl")
	if err != nil {
		return 0, errors.Wrap(err, "os.Create")
	}
	defer func() {
		out.Close()
		utils.RMFile(tmpData + "stdlibdl")
		if c.Debug {
			log.Printf("StdLibDL Download saved to %s", tmpData+"stdlibdl")
		}
	}()

	resp, err := http.Get(c.URL)
	if err != nil {
		return 0, errors.Wrap(err, "http.Get")
	}
	defer resp.Body.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return 0, errors.Wrap(err, "io.Copy")
	}

	return time.Since(t1).Milliseconds(), nil
}

type transfer struct {
	*grab.Response
	ref string
}

func (t transfer) Log(debug bool) {
	if debug {
		charSize := fmt.Sprint(len(fmt.Sprint(t.Size())))
		format := "-%s transferred %0" + charSize + "v / %v bytes - % .2f%%\n"
		fmt.Printf(format, t.ref, t.BytesComplete(), t.Size(), 100*t.Progress())
	}
}

// Sets waiting period in microseconds
func (l limit) WaitN(ctx context.Context, n int) (err error) {
	// log.Printf("%s sleeping for %dµs", l.ref, n) // kept only for debugging purposes
	time.Sleep(time.Duration(n) * time.Microsecond)
	return
}

func (c *Config) Log() {
	log.Printf("c.URL ------- %v", c.URL)
	log.Printf("c.Throttle -- %v", c.Throttle)
	log.Printf("c.Debug ----- %v", c.Debug)
	log.Printf("c.Dest ------ %v", c.Dest)
}
