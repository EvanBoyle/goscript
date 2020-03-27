package goscript

import (
	"errors"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Run executes a "hello" bash script located alongside this file
// the script location is determined dynamically at runtime.
func Run() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("cannot determine location of hello script")
	}

	scriptPath := filepath.Join(filename, "..", "hello")

	cmd := exec.Command("/bin/sh", scriptPath)

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
