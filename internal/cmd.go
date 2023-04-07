package internal

import (
	"log"
	"os/exec"
)

func RunCmd(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}
