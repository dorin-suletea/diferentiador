package data

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type FileStatus string

const (
	added     FileStatus = "a"
	modified  FileStatus = "m"
	deleted   FileStatus = "d"
	untracked FileStatus = "u"
)

func GetChangedFiles() error {
	// split output into lines
	rawGitStatus := runCmd("git", "status", "-s")
	lines := filterEmptyLines(strings.Split(rawGitStatus, "\n"))

	// tokenize into data objects
	files := []File{}
	for _, line := range lines {
		// whitespaces are significant for the status
		statusToken := line[0:2]
		// whitespace are not significant for anything else
		splits := strings.Fields(line)
		switch statusToken {
		case "A ":
			files = append(files, SimpleFile{splits[1], true, added})
		case " A":
			// no-op : unstaged add means untracked
		case "D ":
			files = append(files, SimpleFile{splits[1], true, deleted})
		case " D":
			files = append(files, SimpleFile{splits[1], false, deleted})
		case "M ":
			files = append(files, SimpleFile{splits[1], true, modified})
		case " M":
			files = append(files, SimpleFile{splits[1], false, modified})
		case "R ":
			files = append(files, RenamedFile{SimpleFile{splits[3], true, modified}, splits[1]})
		case " R":
			// no-op : rename is always staged, else its a detele+add
		case "??":
			files = append(files, SimpleFile{splits[1], false, untracked})
		}

	}
	fmt.Printf("%v", files)
	// fmt.Printf("%v", files[1])

	return nil
}

type File interface {
	String() string
}

type SimpleFile struct {
	fileName string
	staged   bool
	status   FileStatus
}

func (f SimpleFile) String() string {
	return fmt.Sprintf("(%v %v %v)", f.fileName, f.staged, f.status)
}

type RenamedFile struct {
	file    SimpleFile
	oldName string
}

func (f RenamedFile) String() string {
	return fmt.Sprintf("(%v -> %v %v %v)", f.oldName, f.file.fileName, f.file.staged, f.file.status)
}

func runCmd(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out[:])
}

// git exec commands have a trailing line at the end, filter it out
func filterEmptyLines(unfiltered []string) []string {
	lines := []string{}
	for _, l := range unfiltered {
		if len(strings.TrimSpace(l)) != 0 {
			lines = append(lines, l)
		}
	}
	return lines
}
