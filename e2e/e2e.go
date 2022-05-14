/*
This is for e2e testing

*/
package e2e

import (
	"errors"
	"os"

	v1 "github.com/cwxstat/dopt/v1"
)

func Setup(testDir string) error {
	if testDir == "" {
		return errors.New("dangerous delete")
	}
	if testDir == "/" {
		return errors.New("dangerous delete")
	}

	err := os.RemoveAll(testDir)
	if err != nil {
		return err
	}
	if err := CreateTestEnv(testDir); err != nil {
		return err
	}
	v1.Init(testDir + "/Dockerfile")
	return nil
}

func createFiles(fileDest string) error {

	err := os.WriteFile(fileDest+"/Dockerfile", dockerfile(), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileDest+"/main.go", goProg(), 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileDest+"/go.mod", goMod(), 0644)
	if err != nil {
		return err
	}
	return nil
}
func CreateTestEnv(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	err = createFiles(dir)
	if err != nil {
		return err
	}
	return nil
}
