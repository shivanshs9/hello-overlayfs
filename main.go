package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
)

const CvmfsPath = "/cvmfs/unpacked.cern.ch"
const MountPath = "/GSoC/unpacked.cern.ch"
const DefaultUpperPath = "./mount"

func main() {
	flag.Parse()
	mntDir := MountPath
	fsInfo, err := NewCvmFs(CvmfsPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		os.Exit(1)
	}
	_ = fsInfo
	err = os.Mkdir(mntDir, os.ModeDir)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(fmt.Errorf("error: %v", err))
		os.Exit(1)
	}
	upperDir := flag.Arg(0)
	if upperDir == "" {
		upperDir = DefaultUpperPath
	}
	opts := fmt.Sprintf("lowerdir=%s:%s", CvmfsPath, upperDir)
	if err := syscall.Mount("overlay", mntDir, "overlay", 0, opts); err != nil {
		fmt.Println(fmt.Errorf("error: failed creating overlay mount to %s: %v", mntDir, err))
	}
}
