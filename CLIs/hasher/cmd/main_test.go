package cmd

import (
	dl "hasher/pkg/download"
	"sync"
	"testing"

	"github.com/jinzhu/copier"
)

func TestBenchMark(t *testing.T) {
	throttled, err := getConfig()
	if err != nil {
		t.Errorf("getConfig - %v", err)
		t.FailNow()
	}

	var (
		throttledTiming   int
		unthrottledTiming int
	)

	unthrottled := dl.Config{}
	copier.Copy(&unthrottled, &throttled)

	unthrottled.Throttle = false
	unthrottled.TmpFile("unthrottled")

	throttled.Throttle = true
	throttled.TmpFile("throttled")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		timestamp, err := unthrottled.GrabPkgDL()
		if err != nil {
			t.Logf("unthrottled.GrabPkgDL - %v", err)
		}
		unthrottledTiming = int(timestamp)
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		timestamp, err := throttled.GrabPkgDL()
		if err != nil {
			t.Logf("  throttled.GrabPkgDL - %v", err)
		}
		throttledTiming = int(timestamp)
		defer wg.Done()
	}()

	wg.Wait()

	t.Logf("")
	t.Logf("unthrottled.GrabPkgDL duration - %dms", unthrottledTiming)
	t.Logf("  throttled.GrabPkgDL duration - %dms", throttledTiming)
	t.Logf("throttled request took %.2f%% more time", (float64(throttledTiming)/float64(unthrottledTiming))*100-100)
	if throttledTiming < unthrottledTiming {
		t.Logf("unthrottled request took longer than throttled request")
		t.Fail()
	}
}
