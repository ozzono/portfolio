package dl

import (
	"fmt"
	"log"
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

func (dl DL) GrabPkgDL() error {
	log.Println("GrabDL")
	t1 := time.Now()
	defer func() {
		log.Printf("GrabDL duration %dms", time.Since(t1).Milliseconds())
	}()

	// create client
	client := grab.NewClient()
	req, err := grab.NewRequest(dl.Dest+"grabdl", dl.URL)
	if err != nil {
		return errors.Wrap(err, "grab.NewRequest")
	}
	defer utils.RMFile(dl.Dest + "grabdl")

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := transfer{client.Do(req)}
	fmt.Printf("- %v\n", resp.HTTPResponse.Status)

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
		return errors.Wrapf(err, "Download failed: %v\n")
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	return nil
}

func (dl DL) GotPkgDL() {
	//under test
	log.Println("GotPkgDL")
	t1 := time.Now()
	defer func() {
		log.Printf("GotPkgDL duration %dms", time.Since(t1).Milliseconds())
	}()

	if err := got.New().Download(dl.URL, dl.Dest); err != nil {
		log.Println(err)
	}
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
