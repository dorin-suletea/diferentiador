package diff

import (
	"github.com/dorin-suletea/diferentiador~/internal"
)

func GetDiffForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "diff", "-U50", filePath)
	return rawGitDiff
}
