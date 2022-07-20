package util

import (
	"os"
	"syscall"
)

// GetExecDirectory Get current executable directory
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}
	return file + "/"
}

// CheckProcessExists will return true if the process with pid exists.
func CheckProcessExists(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}