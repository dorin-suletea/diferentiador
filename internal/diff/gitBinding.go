package diff

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dorin-suletea/diferentiador~/internal"
)

func GetDiffForFile(filePath string) string {
	rawGitDiff := internal.RunCmd("git", "diff", "-U50", filePath)
	return rawGitDiff
}

/*
Useful for dusplaying untracked files.
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
