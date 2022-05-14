/*
This is for e2e testing

*/
package e2e

import (
	"testing"

	"github.com/cwxstat/dopt/tag"
)

func TestCreateTestEnv(t *testing.T) {

	testDir := "./deleteTest/1"
	if Setup(testDir) != nil {
		return
	}
	tag := tag.NewTag()
	err := tag.Read(testDir + "/Dockerfile")
	if err != nil {
		return
	}

	tag.ImageVersion("doptest", "v0.0.1")

}
