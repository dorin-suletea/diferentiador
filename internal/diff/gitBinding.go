package diff

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	// TODO : left off here
	// make the app read the cache instead of doing raw requests.
	return &ret
}

func (gd *GitDifCache) GetContent(key status.FileStatus) string {
	return gd.diffContentMap[key]
}

func (gd *GitDifCache) refresh() {
	for key := range gd.diffContentMap {
		gd.diffContentMap[key] = GetDiffForFile(key.FilePath)
	}
	gd.lastRefreshed = time.Now().Unix()
	fmt.Println(gd.diffContentMap)
}

func GetDiffForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "diff", "-U50", filePath)
	return rawGitDiff
}

/*
Useful for dusplaying untracked files..
`git diff --no-index /dev/null myFilePath` is not portable so will read the file as-is similar to `cat`
*/
func GetRawFileContents(filePath string) string {
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

func GetHeadForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "show", "HEAD^:"+filePath)
	return rawGitDiff
}
