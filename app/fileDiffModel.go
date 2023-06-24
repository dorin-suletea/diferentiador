package app

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dorin-suletea/diferentiador~/internal"
)

type FileDifCache struct {
	// maps each file to it's diff
	diffContentMap map[FileStatus]string
	lastRefreshed  int64
	onRefreshed    func()
}

func NewFileDiffCache(keys []FileStatus, refreshSeconds int) *FileDifCache {
	contentMap := make(map[FileStatus]string)
	for _, fs := range keys {
		contentMap[fs] = ""
	}

	ret := FileDifCache{contentMap, 0, func() { /*no-op*/ }}
	// TODO : this can be done with Promises/channels. Start the cron and block on reader.
	// refresh first blocking and lazily the rest
	if len(keys) != 0 {
		ret.diffContentMap[keys[0]] = ret.invokeGitBindings(keys[0])
	}
	ret.startCron(refreshSeconds)

	return &ret
}

func (gd *FileDifCache) SetOnRefreshHandler(handler func()) {
	gd.onRefreshed = handler
}

func (gd *FileDifCache) GetContent(key FileStatus) string {
	return gd.diffContentMap[key]
}

func (gd *FileDifCache) refresh() {
	for key := range gd.diffContentMap {
		gd.diffContentMap[key] = gd.invokeGitBindings(key)
	}
	gd.lastRefreshed = time.Now().Unix()
	gd.onRefreshed()
}

func (gd *FileDifCache) startCron(refreshSeconds int) {
	go func(refreshSeconds int) {
		gd.refresh()
		for range time.Tick(time.Second * time.Duration(refreshSeconds)) {
			gd.refresh()
		}
	}(refreshSeconds)
}

// Files with different status (modified, deleted, untracked) are issuing different commands for their diff
func (gd *FileDifCache) invokeGitBindings(key FileStatus) string {
	// TODO : fixme. Modified staged dont show
	switch key.Status {
	case Untracked:
		return getRawFileContents(key.FilePath)
	case Modified:
		return getDiffForFile(key.FilePath)
	case Added:
		return markLines(getRawFileContents(key.FilePath), '+')
	case Renamed:
		return key.FilePath
	case Deleted:
		return markLines(getHeadForFile(key.FilePath), '-')

	}
	// TODO handle properly
	panic("Invalid status" + key.Status)
}

func getDiffForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "diff", "-U50", filePath)
	return rawGitDiff
}

// ------------
// Git commands.
// -----------
func getRawFileContents(filePath string) string {
	// Useful for displaying untracked files.
	// `git diff --no-index /dev/null myFilePath` is not portable so will read the file as-is similar to `cat`
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents[:])
}

func getHeadForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "show", "HEAD^:"+filePath)
	return rawGitDiff
}

/*
Prefixed all lines of a given \n separated string.
This is useful to mark all lines of a deleted file with '-' and with '+' for an unstaged file.
While these are not technically line changes from a GIT perspective this provides useful feedback to the user.
*/
func markLines(content string, prefix byte) string {
	prefixed := bytes.Buffer{}
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		prefixed.WriteByte(prefix)
		prefixed.WriteString(line)
		prefixed.WriteString("\n")
	}
	return prefixed.String()
}