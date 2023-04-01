package data

import (
	"fmt"
	"log"
	"os/exec"
)

type FileStatus int32

const (
	untacked FileStatus = iota
	modified
	deleted
	statged
)

func GetChangedFiles() {
	fmt.Printf(runCmd("git", "status"))
}

type File struct {
	FileName string
	staged   bool
}

func runCmd(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
	return string(out[:])
}
