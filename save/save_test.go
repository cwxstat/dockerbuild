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
	err = FileSaveCommit(repo, repoDir, "./save.go", repoDir+"/"+"save.go")
	if err != nil {
		t.FailNow()
	}
	err = os.MkdirAll(repoDir+"/1/2", os.ModePerm)
	if err != nil {
		t.FailNow()
	}
	err = FileSaveCommit(repo, repoDir, "./save.go", repoDir+"/1/2/"+"save.go")
	if err != nil {
		t.FailNow()
	}

}
