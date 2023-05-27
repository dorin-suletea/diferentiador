package diff

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dorin-suletea/diferentiador~/internal"
	"github.com/dorin-suletea/diferentiador~/internal/status"
)

type GitDifCache struct {
	// maps each file to it's diff
	diffContentMap map[status.FileStatus]string
	lastRefreshed  int64
}

func NewGitDiffCache(keys []status.FileStatus) *GitDifCache {
	contentMap := make(map[status.FileStatus]string)

	for _, fs := range keys {
		contentMap[fs] = ""
	}

	ret := GitDifCache{contentMap, 0}
	ret.refresh()
	return &ret
}

func (gd *GitDifCache) GetContent(key status.FileStatus) string {
	// todo : precompute getRawFileContents| getHeadForFile etc in init
	// switch key.Status {
	// case status.Deleted:
	// 	content := markLines(getHeadForFile(key.FilePath), '-')
	// 	return content
	// case status.Untracked:
	// 	content := getRawFileContents(key.FilePath)
	// 	return content
	// case status.Modified:
	// 	content := getDiffForFile(key.FilePath)
	// 	return content
	// }

	return gd.diffContentMap[key]
}

func (gd *GitDifCache) refresh() {
	for key := range gd.diffContentMap {
		// files with different status (modified, deleted, untracked) are issuing different commands for their diff
		switch key.Status {
		case status.Deleted:
			gd.diffContentMap[key] = markLines(getHeadForFile(key.FilePath), '-')
		case status.Untracked:
			gd.diffContentMap[key] = getRawFileContents(key.FilePath)
		case status.Added:
			gd.diffContentMap[key] = markLines(getRawFileContents(key.FilePath), '+')
		case status.Modified:
			gd.diffContentMap[key] = getDiffForFile(key.FilePath)
		}
	}
	gd.lastRefreshed = time.Now().Unix()
	fmt.Println(gd.diffContentMap)
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
