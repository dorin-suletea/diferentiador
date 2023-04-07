package diff

import (
	"github.com/dorin-suletea/diferentiador~/internal"
)

func GetDiffForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "diff", "-U100", "main.go")
	return rawGitDiff
}
