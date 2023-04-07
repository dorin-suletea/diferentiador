package diff

import (
	"log"

	"github.com/dorin-suletea/diferentiador~/internal"
)

func GetDiffForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "diff", "-U100", "main.go")
	log.Println(rawGitDiff)
	return rawGitDiff
}
