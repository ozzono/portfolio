package utils

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	sampleSrcFiles = []string{
		"%s/1/1.1/1.1.1/1.1.1.txt",
		"%s/1/1.1/1.1.1/1.1.2.txt",
		"%s/1/1.1/1.1.txt",
		"%s/1/1.txt",
		"%s/2/2.txt",
	}
	sampleSrcFolders = []string{
		"%s/1",
		"%s/1/1.1",
		"%s/1/1.1/1.1.1",
		"%s/2",
	}
	sampleDestFiles = []string{
		"%s/1/1.1/1.1.1/1.1.1.txt",
		"%s/1/1.1/1.1.1/1.1.2.txt",
		"%s/1/1.1/1.1.txt",
		"%s/1/1.txt",
		"%s/2/2.txt",
	}
	sampleDestFolders = []string{
		"%s/1",
		"%s/1/1.1",
		"%s/1/1.1/1.1.1",
		"%s/2",
	}
	sampleSrc   string
	sampleDest  string
	sampleFiles PathData
)

func init() {
	flag.StringVar(&sampleSrc, "sample-src", "./sample", "Sample sourcedata path")
	flag.StringVar(&sampleDest, "sample-dest", "./destiny", "Sample destiny data path")
}

func TestListPath(t *testing.T) {
	if err := validateArgs(t); err != nil {
		t.Errorf("impossible to perfom test")
		t.Error(err)
		t.FailNow()
	} else {
		t.Logf("testing sample path - %s", sampleSrc)
	}
	pathData, err := ListPath(sampleSrc, true)
	if err != nil {
		t.Errorf("ListPath - %v", err)
	}
	sampleFiles = *pathData

	files := toMap(pathData.Files)
	for i := range sampleSrcFiles {
		_, ok := files[sampleSrcFiles[i]]
		if !ok {
			t.Errorf("missing sample file: %v", sampleSrcFiles[i])
		}
	}

	folders := toMap(pathData.Folders)
	for i := range sampleSrcFolders {
		_, ok := folders[sampleSrcFolders[i]]
		if !ok {
			t.Errorf("missing sample folder: %v", sampleSrcFolders[i])
		}
	}
}

func TestMkdirAll(t *testing.T) {
	if err := sampleFiles.MkdirAll(sampleSrc, sampleDest); err != nil {
		t.Errorf("failed to make files - %v", err)
	}
}

func TestCpAll(t *testing.T) {
	if err := sampleFiles.CpAll(sampleSrc, sampleDest); err != nil {
		t.Errorf("failed to cp files - %v", err)
	}
}

func TestCopiedFiles(t *testing.T) {
	pathData, err := ListPath(sampleDest, true)
	if err != nil {
		t.Errorf("ListPath - %v", err)
		t.FailNow()
	}
	sampleFiles = *pathData

	files := toMap(pathData.Files)
	for i := range sampleDestFiles {
		_, ok := files[sampleDestFiles[i]]
		if !ok {
			t.Errorf("missing sample file: %v", sampleDestFiles[i])
		}
	}

	folders := toMap(pathData.Folders)
	for i := range sampleDestFolders {
		_, ok := folders[sampleDestFolders[i]]
		if !ok {
			t.Errorf("missing sample folder: %v", sampleDestFolders[i])
		}
	}
}

func TestEraseTestTrack(t *testing.T) {
	if err := os.RemoveAll(sampleDest); err != nil {
		t.Errorf("os.RemoveAll - %v", err)
		t.Fail()
	}
}

func toMap(f []string) map[string]string {
	output := map[string]string{}
	for i := range f {
		output[f[i]] = f[i]
	}
	return output
}

func validateArgs(t *testing.T) error {
	if sampleSrc == "" {
		return fmt.Errorf("invalid sample source path; cannot be nil")
	}
	if sampleDest == "" {
		return fmt.Errorf("invalid sample destiny path; cannot be nil")
	}

	src := []string{}
	for i := range sampleSrcFiles {
		src = append(src, fmt.Sprintf(sampleSrcFiles[i], sampleSrc))
	}
	t.Log("setting sample source files")
	sampleSrcFiles = src

	srcFolder := []string{}
	for i := range sampleSrcFolders {
		srcFolder = append(srcFolder, fmt.Sprintf(sampleSrcFolders[i], sampleSrc))
	}
	t.Log("setting sample source folders")
	sampleSrcFolders = srcFolder

	dest := []string{}
	for i := range sampleDestFiles {
		dest = append(dest, fmt.Sprintf(sampleDestFiles[i], sampleDest))
	}
	t.Log("setting sample destiny files")
	sampleDestFiles = dest

	destFolders := []string{}
	for i := range sampleDestFolders {
		destFolders = append(destFolders, fmt.Sprintf(sampleDestFolders[i], sampleDest))
	}
	t.Log("setting sample destiny folders")
	sampleDestFolders = destFolders

	return nil
}
