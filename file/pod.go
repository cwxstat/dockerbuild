package file

import (
	"gopkg.in/yaml.v2"
)

type Pod struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			AppKubernetesIoName string `yaml:"app.kubernetes.io/name"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Volumes []struct {
			Name                  string `yaml:"name"`
			PersistentVolumeClaim struct {
				ClaimName string `yaml:"claimName"`
				ReadOnly  bool   `yaml:"readOnly"`
			} `yaml:"persistentVolumeClaim"`
		} `yaml:"volumes"`
		Containers []struct {
			Name  string `yaml:"name"`
			Image string `yaml:"image"`
			Env   []struct {
				Name      string `yaml:"name"`
				ValueFrom struct {
					SecretKeyRef struct {
						Name string `yaml:"name"`
						Key  string `yaml:"key"`
					} `yaml:"secretKeyRef"`
				} `yaml:"valueFrom"`
			} `yaml:"env"`
			VolumeMounts []struct {
				MountPath string `yaml:"mountPath"`
				Name      string `yaml:"name"`
			} `yaml:"volumeMounts"`
		} `yaml:"containers"`
		RestartPolicy string `yaml:"restartPolicy"`
	} `yaml:"spec"`
}

// Ref: https://go.dev/play/p/NTM3D2NGkry

func FindImages(data string) ([]string, error) {
	var p Pod
	out := []string{}
	err := yaml.Unmarshal([]byte(data), &p)
	if err != nil {
		return out, err
	}
	for _, v := range p.Spec.Containers {
		out = append(out, v.Image)
	}
	return out, nil
}
