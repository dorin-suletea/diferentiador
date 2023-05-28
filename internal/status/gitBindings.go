package status

import (
	"strings"

	"github.com/dorin-suletea/diferentiador~/internal"
)

type Status string

const (
	Added     Status = "a"
	Modified  Status = "m"
	Deleted   Status = "d"
	Renamed   Status = "r"
	Untracked Status = "u"
)

func GetStatusForFiles() []FileStatus {
	// alternatively `internal.RunCmd("git", "status", "-s", "-u")``
	rawGitStatus := internal.RunCmd("git", "status", "--porcelain")
	lines := filterEmptyLines(strings.Split(rawGitStatus, "\n"))

	// TODO : there are many more status

	// tokenize into data objects
	files := []FileStatus{}
	for _, line := range lines {
		// whitespaces are significant for the stat us
		statusToken := line[0:2]
		// whitespace are not significant for anything else
		splits := strings.Fields(line)
		switch statusToken {
		case "A ":
			files = append(files, FileStatus{splits[1], true, Added})
		case " A":
			// no-op : unstaged add means untracked
		case "D ":
			files = append(files, FileStatus{splits[1], true, Deleted})
		case " D":
			files = append(files, FileStatus{splits[1], false, Deleted})
		case "M ":
			files = append(files, FileStatus{splits[1], true, Modified})
		case " M":
			files = append(files, FileStatus{splits[1], false, Modified})
		case "MM":
			files = append(files, FileStatus{splits[1], false, Modified})
		case "R ":
			oldName := splits[3]
			newName := splits[1]
			files = append(files, FileStatus{oldName + "->" + newName, true, Renamed})
		case " R":
			// no-op : rename is always staged, else its a detele+add
		case "??":
			files = append(files, FileStatus{splits[1], false, Untracked})
		default:
			panic("Fixme : Can't parse : " + statusToken)
		}
	}
	return files
}

type FileStatus struct {
	FilePath string
	staged   bool
	Status   Status
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
