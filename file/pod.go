package file

import (
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kubernetes/pkg/apis/apps"
)

func Pod(filename string) (*v1.Pod, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pod := &v1.Pod{}
	if err := yaml.NewYAMLOrJSONDecoder(f, 4096).Decode(pod); err != nil {
		return nil, err
	}

	return pod, nil

}

func Deployment(filename string) (*apps.Deployment, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	deployment := &apps.Deployment{}
	if err := yaml.NewYAMLOrJSONDecoder(f, 4096).Decode(deployment); err != nil {
		return nil, err
	}

	return deployment, nil

}
