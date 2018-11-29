package main

// scratch
// Disposable command line notes

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scratchpath() string {
	usr, err := user.Current()
	check(err)
	// Remove .swp while we're at it
	// TODO: Pull into more explicit function
	swp := filepath.Join(usr.HomeDir, ".scratchpad.md.swp")
	cmd := exec.Command("rm", swp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	check(err)
	return filepath.Join(usr.HomeDir, "scratchpad.md")
}

func makePad(p string) {
	f, err := os.Create(p)
	check(err)
	defer f.Close()
	_, err = f.WriteString("# Scratchpad\n\n\n")
	f.Sync()
	check(err)
}

func openPad(p string) {
	cmd := exec.Command("vim", p)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	check(err)
}

func scratch() {
	p := scratchpath()
	makePad(p)
	openPad(p)
}

func main() {
	scratch()
}
