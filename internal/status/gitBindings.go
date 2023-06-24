package status

import (
	"fmt"
	"strings"
	"time"

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

func getStatusForFiles() []FileStatus {
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

type FileStatusCache struct {
	status        []FileStatus
	lastRefreshed int64
	onRefreshed   func()

	//TODO impl generic promise

	//This allows marginally better performance.
	//The app needs initial file statuses to continue with other initializations.
	//Instead of issuing a blocking request to git on cache creation
	//we will do this in a gorouting, allow the flow of the app to continue with other stuff
	//and block only when the data is actually needed. Think a Promise/Future.
	boostrapDone bool
	boostrapChan chan []FileStatus
}

func NewChangedFilesCache(refreshSeconds int) *FileStatusCache {
	ret := &FileStatusCache{[]FileStatus{}, 0, func() { /*no op*/ }, false, make(chan []FileStatus)}

	go func(bc chan []FileStatus) {
		bc <- getStatusForFiles()
	}(ret.boostrapChan)

	ret.startCron(refreshSeconds)
	fmt.Println("exit new")
	return ret
}

func (t *FileStatusCache) GetChangedFiles() []FileStatus {
	fmt.Println("enter get")
	if !t.boostrapDone {
		fmt.Println("blocking")
		t.status = <-t.boostrapChan
		t.boostrapDone = true
		fmt.Println("exiting")
	}

	return t.status
}

func (t *FileStatusCache) SetOnRefreshHandler(handler func()) {
	t.onRefreshed = handler
}

func (t *FileStatusCache) startCron(refreshSeconds int) {
	go func(refreshSeconds int, gdLocal *FileStatusCache) {
		for range time.Tick(time.Second * time.Duration(refreshSeconds)) {
			gdLocal.refresh()
		}
	}(refreshSeconds, t)
}

func (t *FileStatusCache) refresh() {
	t.status = getStatusForFiles()
	t.lastRefreshed = time.Now().Unix()
	t.onRefreshed()
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
