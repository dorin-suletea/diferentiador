package internal

import (
	"log"
	"os/exec"
)

// TODO : Figure out if and how to run in a foreign directory.
func RunGit(rootPath string, args ...string) string {
	cmd := "git"
	pa := append([]string{"-C", rootPath}, args...)

	out, err := exec.Command(cmd, pa...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}

func RunCmd(cmd string, args ...string) string {
	// TODO : bug it will try to run "HEAD^:internal/fileDiffModel.go"
	// altho the file was moved but not commited into app/fileDiffModel.go
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}
