package cmd

import (
	dl "hasher/pkg/download"
	"sync"
	"testing"

	"github.com/jinzhu/copier"
)

func TestThrottleBench(t *testing.T) {
	throttled, err := getConfig()
	if err != nil {
		t.Errorf("getConfig - %v", err)
		t.FailNow()
	}

	unthrottled := dl.Config{}
	copier.Copy(&unthrottled, &throttled)

	unthrottled.Throttle = false
	unthrottled.TmpFile("unthrottled")

	throttled.Throttle = true
	throttled.TmpFile("throttled")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		if _, err := unthrottled.GrabPkgDL(); err != nil {
			t.Logf("unthrottled.GrabPkgDL - %v", err)
		}
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		if _, err := throttled.GrabPkgDL(); err != nil {
			t.Logf("  throttled.GrabPkgDL - %v", err)
		}
		defer wg.Done()
	}()

	wg.Wait()

	t.Logf("unthrottled.GrabPkgDL duration -- %dms", unthrottled.ElapsedTime)
	t.Logf("throttled.GrabPkgDL duration ---- %dms", throttled.ElapsedTime)
	t.Logf("throttled request took %.2f%% more time", (float64(throttled.ElapsedTime)/float64(unthrottled.ElapsedTime))*100-100)
	if throttled.ElapsedTime < unthrottled.ElapsedTime {
		t.Logf("unthrottled request took longer than throttled request")
		t.Fail()
	}
}
