/*
This is for e2e testing

*/
package e2e

import (
	"testing"

	"os"
	"strings"

	"github.com/cwxstat/dopt/bcmds"
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
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	cmd := "buildx build --no-cache --progress=plain --platform linux/amd64 --no-cache -t us-central1-docker.pkg.dev/mchirico/public/dopt:v0.0.1 -f Dockerfile ."
	scmd := strings.Fields(cmd)
	bcmds.DockerDir(currentDir+"/deleteTest/1", scmd...)

}
