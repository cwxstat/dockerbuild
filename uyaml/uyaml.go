package uyaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type FilesUpdate struct {
	Files []string `yaml:"files,omitempty"`
}

type Platform struct {
	OS string `yaml:"os,omitempty"`
}

type Spec struct {
	Timestamp time.Time `yaml:"timestamp,omitempty"`
	Platform         Platform        `yaml:"platform,omitempty"`
	Files FilesUpdate `yaml:"changes,omitempty"`
}

type T struct {
	Image   string `yaml:"image"`
	Version string `yaml:"version"`
	Config string `yaml:"local-config,omitempty"`
	Spec    Spec   `yaml:"spec,omitempty"`
}

func MyTest() {

	fu := FilesUpdate{
		Files: []string{"pod.yaml, dev-pod.yaml"},
	}
	t := &T{
		Image:   "ubuntu",
		Version: "v0.0.1",
		Config: "~/.docTag/config",
		Spec: Spec{
			Timestamp: time.Now(),
			Platform: Platform{
				OS: "linux/amd64",
			},
			Files: fu,
		},
	}

	if a, err := yaml.Marshal(t); err == nil {
		fmt.Println(a)
		os.WriteFile("sample.yaml", []byte(a), 0644)
	}

}
