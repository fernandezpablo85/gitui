package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if !isGitInstalled() {
		exitOne("gitui requires git installed")
	}
	if !isGitRoot() {
		exitOne("gitui must be invoked at a git repository root")
	}
}

func isGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()
	return err == nil
}

func isGitRoot() bool {
	wd, err := os.Getwd()

	if err != nil {
		panic("cant get current directory")
	}

	_, err2 := os.Stat(wd + "/.git")
	return err2 == nil
}

func exitOne(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
