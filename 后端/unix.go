// +build darwin freebsd linux

package main
func GetLogicalDrives() []fileInfo{
	return getFile("/")
}

