package file

import (
	"testing"
)

func TestFindImages(t *testing.T) {
	CreatePodDeploymentFiles()
	pod, err := Pod("./testPod.yaml")
	if err != nil {
		t.Fatalf("pod: %v", err)
	}
	if pod.Name != "nginx" {
		t.FailNow()
	}

	_, err = Pod("./testDeployment.yaml")
	if err != nil {
		t.Fatalf("deployment: %v", err)
	}

}
