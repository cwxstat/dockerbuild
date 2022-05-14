package save

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	repoDir := "/tmp/dopt/test"
	os.RemoveAll(repoDir)
	repo, err := Init(repoDir)
	if err != nil {
		t.FailNow()
	}
	err = FileSaveCommit(repo, repoDir, "./save.go")
	if err != nil {
		t.FailNow()
	}

	err = FileSaveCommit(repo, repoDir, "../constants/constants.go")
	if err != nil {
		t.FailNow()
	}

}
