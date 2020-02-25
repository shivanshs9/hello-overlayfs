package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/afero"
)

type CvmFs struct {
	source afero.Fs
	Path   string
}

type InvalidFs struct {
	path string
}

func (e *InvalidFs) Error() string { return fmt.Sprintf("'%s' is not a CvmFS filesystem", e.path) }

const FsName = "cvmfs"

func NewCvmFs(path string) (fs *CvmFs, err error) {
	// Use linux 'df' command to check filesystem of target path
	// Can read /proc/mounts for a POSIX-way to do the same.
	cmd := exec.Command("df", "--output=source", path)
	out, err := cmd.Output()
	if err != nil {
		return fs, err
	}
	if !strings.Contains(string(out), FsName) {
		err = &InvalidFs{path: path}
		return fs, err
	}
	source := afero.NewOsFs()
	fs = &CvmFs{source: source, Path: path}
	return fs, err
}

func (CvmFs) Name() string {
	return FsName
}
