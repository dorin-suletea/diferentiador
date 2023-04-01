package data

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type FileStatus int32

const (
	added    FileStatus = 0
	modified FileStatus = 1
	deleted  FileStatus = 2
)

func GetChangedFiles() error {
	files := []File{}

	deletedFilesStr := runCmd("git", "status", "-s")
	lines := splitByEmptyNewline(deletedFilesStr)
	for _, line := range lines {
		splits := strings.Split(line, " ")
		if len(splits) > 2 {
			return errors.New("InternalError line must have 2 tokens :" + line)
		}
		switch splits[0] {
		case "A":
			files = append(files, File{splits[1], true, added})
		case "D":
			files = append(files, File{splits[1], true, deleted})
		case "M":
			files = append(files, File{splits[1], true, modified})
		}
	}

	fmt.Printf("%v", lines)
}

// only deleted
//git ls-files -d

// modified and deleted
//git ls-files -m

// all staged
//git ls-files -s

// all unstaged
//git ls-files -o

type File struct {
	FileName string
	staged   bool
	status   FileStatus
}

func runCmd(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}
