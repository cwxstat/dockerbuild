/*
This is for e2e testing

*/
package e2e

import (
	"testing"
)

func TestCreateTestEnv(t *testing.T) {

	testDir := "./deleteTest/1"
	if Setup(testDir) != nil {
		return
	}

}
