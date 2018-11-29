package main

// scratch
// Disposable command line notes

import (
	"fmt"
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
	f := filepath.Join(usr.HomeDir, "scratchpad.md")
	if exists(swp) {
		fmt.Println(".scratchpad.md.swp contents:\n")
		cat(swp)
		cmd := exec.Command("rm", swp)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		check(err)
	}
	fmt.Println("scratchpad.md contents:\n")
	cat(f)
	return f
}

func makePad(p string) {
	f, err := os.Create(p)
	check(err)
	defer f.Close()
	_, err = f.WriteString("# Scratchpad\n\n\n")
	f.Sync()
	check(err)
}

func cat(p string) {
	cmd := exec.Command("cat", p)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
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

func exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func scratch() {
	p := scratchpath()
	makePad(p)
	openPad(p)
}

func main() {
	scratch()
}
