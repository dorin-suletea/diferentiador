package internal

import (
	"log"
	"os/exec"
)

func Git(workdir string, args ...string) string {
	cmd := "git"
	pa := append([]string{"-C", workdir}, args...)

	out, err := exec.Command(cmd, pa...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}

func runCmd(cmd string, args ...string) string {
	// TODO : bug it will try to run "HEAD^:internal/fileDiffModel.go"
	// altho the file was moved but not commited into app/fileDiffModel.go
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}
