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

	pod, err = Pod("./testDeployment.yaml")
	if err != nil {
		t.Fatalf("deployment: %v", err)
	}

	// Note: careful
	if pod.TypeMeta.Kind != "Deployment" {
		t.FailNow()
	}

	if pod.Name != "nginx-deployment" {
		t.Fatalf("pod name from deployment")
	}

}
