package manifest

import (
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"io/ioutil"
	"strings"
)

// Replace '"RepoTags":null' string
// TODO Reading and Parsing JSON file to struct
func ModifyRepoTags(file, tag string) {
	oldStr := "\"RepoTags\":null"
	newStr := "\"RepoTags\":"+"[\""+tag+"\"]"

	read, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Fatal("Cannot read manifest file '%s' with error: '%s'", file, err)
	}

	newContents := strings.Replace(string(read), oldStr, newStr, -1)

	err = ioutil.WriteFile(file, []byte(newContents), 0)
	if err != nil {
		logger.Fatal("Cannot write manifest file '%s' with error: '%s'", file, err)
	}

	fs.EraseDataCreation(file)
}
