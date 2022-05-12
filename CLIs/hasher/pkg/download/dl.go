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

type DL struct {
	Dest     string
	URL      string
	Throttle bool
	Debug    bool
}

type limit struct{}

func (dl DL) GrabPkgDL() (int64, error) {
	log.Println("GrabDL")
	t1 := time.Now()

	// create client
	client := grab.NewClient()

	req, err := grab.NewRequest(dl.Dest+"grabdl", dl.URL)
	if err != nil {
		return 0, errors.Wrap(err, "grab.NewRequest")
	}
	if dl.Throttle {
		req.RateLimiter = limit{}
		if err := req.RateLimiter.WaitN(req.Context(), 1000); err != nil {
			log.Fatalf("req.RateLimiter.WaitN - %v", err)
		}
	}
	defer utils.RMFile(dl.Dest + "grabdl")

	// start download
	if dl.Debug {
		log.Printf("GrabDL Downloading %v", req.URL())
	}

	resp := transfer{client.Do(req)}
	if dl.Debug {
		log.Printf("- %v", resp.HTTPResponse.Status)
	}

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			resp.Log(dl.Debug)
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		return 0, errors.Wrapf(err, "GrabDL Download failed: %v\n")
	}

	return time.Since(t1).Milliseconds(), nil
}

func (dl DL) GotPkgDL() (int64, error) {
	log.Println("GotPkgDL")
	t1 := time.Now()

	if err := got.New().Download(dl.URL, dl.Dest+"gotpkgdl"); err != nil {
		return 0, errors.Wrap(err, "got.New().Download()")
	}

	return time.Since(t1).Milliseconds(), nil
}

func (dl DL) StdLibDL() (int64, error) {
	log.Println("StdLibDL")

	t1 := time.Now()
	out, err := os.Create(dl.Dest + "stdlibdl")
	if err != nil {
		return 0, errors.Wrap(err, "os.Create")
	}
	defer func() {
		out.Close()
		utils.RMFile(dl.Dest + "stdlibdl")
		if dl.Debug {
			log.Printf("StdLibDL Download saved to %s", dl.Dest+"stdlibdl")
		}
	}()

	resp, err := http.Get(dl.URL)
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
}

func (t transfer) Log(debug bool) {
	if debug {
		charSize := fmt.Sprint(len(fmt.Sprint(t.Size())))
		format := "- transferred %0" + charSize + "v / %v bytes - % .2f%%\n"
		fmt.Printf(format, t.BytesComplete(), t.Size(), 100*t.Progress())
	}
}

func (l limit) WaitN(ctx context.Context, n int) (err error) {
	time.Sleep(time.Duration(n) * time.Microsecond)
	return nil
}
