package goscript

import (
	"io/ioutil"
	"os/exec"
)

// Run writes a "hello" script to a temp file and executes it
func Run() (string, error) {
	f, err := ioutil.TempFile("", "hello")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(script)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("/bin/sh", f.Name())

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

const script = `
#!/bin/bash
echo "hello_world"
`
