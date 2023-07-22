package app

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dorin-suletea/diferentiador~/internal"
)

type FileDifCache struct {
	diffContentMap   map[FileStatus]string
	lastRefreshed    int64
	refreshListeners []internal.CacheListener
	fscache          *ChangedFileCache
	gitWorkDir       string
}

func DiffCache(fsCache *ChangedFileCache, workDir string, refreshSeconds int) *FileDifCache {
	ret := FileDifCache{nil, 0, []internal.CacheListener{}, fsCache, workDir}
	ret.startCron(refreshSeconds)
	return &ret
}

func (t *FileDifCache) RegisterCacheListener(listener internal.CacheListener) {
	t.refreshListeners = append(t.refreshListeners, listener)
}

func (gd *FileDifCache) GetContent(key FileStatus) string {
	return gd.diffContentMap[key]
}

func (t *FileDifCache) startCron(refreshSeconds int) {
	go func(refreshSeconds int, tLocal *FileDifCache) {
		//a) eagerly refresh as soon as possible
		t.refresh()
		//b) Keep refreshing the cache on a cron in case new files are added or removed.
		for range time.Tick(time.Second * time.Duration(refreshSeconds)) {
			t.refresh()
		}
	}(refreshSeconds, t)
}

func (t *FileDifCache) refresh() {
	keys := t.fscache.GetAll()
	content := make(map[FileStatus]string, len(keys))
	for _, key := range keys {
		diff, err := runGitDiff(t.gitWorkDir, key)
		if err != nil {
			fmt.Println(err)
		} else {
			content[key] = diff
		}
	}
	t.diffContentMap = content
	t.lastRefreshed = time.Now().Unix()

	for i := range t.refreshListeners {
		t.refreshListeners[i].OnCacheRefreshed()
	}
}

// Files with different statuses(modified, deleted, untracked etc) need different approaches to extracting their content.
func runGitDiff(workdir string, key FileStatus) (string, error) {
	// TODO : fixme. Modified staged dont show
	switch key.Status {
	case Untracked:
		return readFileDirectly(workdir, key), nil
	case Modified:
		return internal.Git(workdir, "diff", "-U50", key.FilePath), nil
	case Added:
		return prefixAllLines(readFileDirectly(workdir, key), '+'), nil
	case Renamed:
		return key.FilePath, nil
	case Deleted:
		return prefixAllLines(getHeadForFile(workdir, key), '-'), nil
	}
	return "", errors.New("Unknown status" + string(key.Status))
}

// Mostly for displaying untracked files.
// `git diff --no-index /dev/null myFilePath` is not portable hence we read the content as-is (similar to `cat`)
func readFileDirectly(workdir string, key FileStatus) string {
	relPath := filepath.Join(workdir, key.FilePath)
	contents, err := os.ReadFile(relPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents[:])
}

func getHeadForFile(workdir string, key FileStatus) string {
	rawGitDiff := internal.Git(workdir, "show", "HEAD^:"+key.FilePath)
	return rawGitDiff
}

// Prefix all the lines of a new file with `+` (respectively `-` for deleted files).
// This makes them show as added lines, abeit they are not new lines from git's perspective, the entire file is new.
func prefixAllLines(content string, prefix byte) string {
	prefixed := bytes.Buffer{}
	for _, line := range strings.Split(strings.TrimSuffix(content, "\n"), "\n") {
		prefixed.WriteByte(prefix)
		prefixed.WriteString(line)
		prefixed.WriteString("\n")
	}
	return prefixed.String()
}
