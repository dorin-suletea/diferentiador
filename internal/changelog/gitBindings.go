package changelog

import (
	"log"
	"os/exec"
	"strings"
)

type FileStatus string

const (
	Added     FileStatus = "a"
	Modified  FileStatus = "m"
	Deleted   FileStatus = "d"
	Untracked FileStatus = "u"
)

func GetGitModififications() []GitModification {
	// split output into lines
	rawGitStatus := runCmd("git", "status", "-s", "-u")
	lines := filterEmptyLines(strings.Split(rawGitStatus, "\n"))

	// tokenize into data objects
	files := []GitModification{}
	for _, line := range lines {
		// whitespaces are significant for the status
		statusToken := line[0:2]
		// whitespace are not significant for anything else
		splits := strings.Fields(line)
		switch statusToken {
		case "A ":
			files = append(files, GitModification{splits[1], true, Added})
		case " A":
			// no-op : unstaged add means untracked
		case "D ":
			files = append(files, GitModification{splits[1], true, Deleted})
		case " D":
			files = append(files, GitModification{splits[1], false, Deleted})
		case "M ":
			files = append(files, GitModification{splits[1], true, Modified})
		case " M":
			files = append(files, GitModification{splits[1], false, Modified})
		case "R ":
			oldName := splits[3]
			newName := splits[1]
			files = append(files, GitModification{oldName + "->" + newName, true, Modified})
		case " R":
			// no-op : rename is always staged, else its a detele+add
		case "??":
			files = append(files, GitModification{splits[1], false, Untracked})
		}
	}
	return files
}

type GitModification struct {
	fileName string
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
