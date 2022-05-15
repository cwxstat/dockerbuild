/*
This is for e2e testing

*/
package e2e

import (
	"testing"

	"os"

	"github.com/cwxstat/dopt/dockersdk"
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

	image := "us-central1-docker.pkg.dev/mchirico/public/dopt"
	version := "v0.0.1"
	tag.ImageVersion(image, version)
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	d, err := dockersdk.NewDocker()
	if err != nil {
		return
	}
	d.Tar(currentDir + "/deleteTest/1")
	d.Image(image)
	d.Version(version)
	d.Platform("linux/amd64")
	err = d.ImageBuild()
	if err != nil {
		t.Errorf("can't build image")
	}
	d.Push()

}
