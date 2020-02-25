# Hello OverlayFS

## Description
A small GO utility that uses [OverlayFS](https://wiki.archlinux.org/index.php/Overlay_filesystem) to overlay the `./mount/` directory on top of CVMFS filesystem mounted path `/cvmfs/unpacked.cern.ch` and the merged filesystem is mounted at `/GSoC/unpacked.cern.ch` by default. The `./mount/` contains `gsoc/print_hello` binary to just print "Hello" on screen when invoked.

## How to run?
- Make sure CVMFS filesystem is mounted at `/cvmfs/unpacked.cern.ch`.
- Make sure `/GSoC/unpacked.cern.ch` directory exists.
- Run the program in **sudo mode**:
```shell
sudo go run *.go
```