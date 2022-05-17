package utils

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors" // allow me the poetic license to use only this external package
)

type PathData struct {
	Folders []string
	Files   []string
	Verbose bool
}

func ListPath(path string, verbose bool) (*PathData, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadDir")
	}

	output := PathData{Verbose: verbose}
	for _, file := range files {
		fPath := fmt.Sprintf("%s/%s", path, file.Name())
		ftype := "file"
		if file.IsDir() {
			ftype = "folder"
			output.Folders = append(output.Folders, fPath)
			subfile, err := ListPath(fPath, verbose)
			if err != nil {
				return nil, errors.Wrap(err, "listPath(fPath)")
			}
			output.Files = append(output.Files, subfile.Files...)
			output.Folders = append(output.Folders, subfile.Folders...)
			continue
		}
		output.Files = append(output.Files, fPath)

		if verbose {
			fmt.Printf("% 6s %s/%s\n", ftype, path, file.Name())
		}
	}
	return &output, nil
}

func (path PathData) MkdirAll(source, destiny string) error {
	var e error
	for i := range path.Folders {
		d := strings.TrimPrefix(path.Folders[i], source)
		d = fmt.Sprintf("%s%s", destiny, d)
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			e = errors.Wrapf(e, fmt.Sprintf("\n%d - %v os.MkdirAll", i, err))
		}
		if path.Verbose {
			fmt.Printf("creating folder %s\n", d)
		}
	}
	return e
}

func (path PathData) CpAll(source, destiny string) error {

	for i := range path.Files {
		d := strings.TrimPrefix(path.Files[i], source)
		d = fmt.Sprintf("%s%s", destiny, d)
		copy(path.Files[i], d, path.Verbose)
	}
	return nil
}

func copy(src, dest string, verbose bool) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return errors.Wrap(err, "ioutil.ReadFile")
	}

	if err = ioutil.WriteFile(dest, input, fs.ModePerm); err != nil {
		return errors.Wrap(err, "ioutil.WriteFile")
	}

	if verbose {
		pattern := `from %s
to   %s
`
		fmt.Printf(pattern, src, dest)
	}
	return nil
}
