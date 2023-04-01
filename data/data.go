package data

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type FileStatus string

const (
	added    FileStatus = "a"
	modified FileStatus = "m"
	deleted  FileStatus = "d"
	unstaged FileStatus = "u"
)

func GetChangedFiles() error {
	// split output into lines
	deletedFilesStr := runCmd("git", "status", "-s")
	unfilteredLines := strings.Split(deletedFilesStr, "\n")
	fmt.Println("----------")
	lines := []string{}
	for _, l := range unfilteredLines {
		if len(strings.TrimSpace(l)) != 0 {
			lines = append(lines, l)
		}
	}

	fmt.Println(lines)
	fmt.Println("----------")

	// tokenize into data objects
	files := []File{}
	for _, line := range lines {
		splits := strings.Fields(line)
		// fmt.Println(line)
		// fmt.Println("----------")
		if len(splits) > 2 {
			fmt.Println("InternalError line must have 2 tokens :" + line)
		}
		switch splits[0] {
		case "A":
			files = append(files, File{splits[1], true, added})
		case "D":
			files = append(files, File{splits[1], true, deleted})
		case "M":
			files = append(files, File{splits[1], true, modified})
		case "MM":
			files = append(files, File{splits[1], true, modified})
		case "??":
			files = append(files, File{splits[1], true, unstaged})
		}

	}
	fmt.Printf("%v", files)
	// fmt.Printf("%v", files[1])

	return nil
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
