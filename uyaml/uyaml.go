package uyaml

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type FilesUpdate struct {
	Files []string `yaml:"files,omitempty"`
}

type Platform struct {
	OS string `yaml:"os,omitempty"`
}

type Spec struct {
	Timestamp time.Time   `yaml:"timestamp,omitempty"`
	Platform  Platform    `yaml:"platform,omitempty"`
	Files     FilesUpdate `yaml:"changes,omitempty"`
}

type TopYaml struct {
	Image   string `yaml:"image"`
	Version string `yaml:"version"`
	Config  string `yaml:"local-config,omitempty"`
	Spec    Spec   `yaml:"spec,omitempty"`
}

func MyTest() string {

	fu := FilesUpdate{
		Files: []string{"pod.yaml, dev-pod.yaml"},
	}
	t := &TopYaml{
		Image:   "ubuntu",
		Version: "v0.0.1",
		Config:  "~/.docTag/config",
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
		return string(a)
	}
	return ""
}

func AddComments(s string) string {
	split := strings.Split(s, "\n")
	newString := ""
	sep := ""
	for _, v := range split {
		if strings.HasPrefix(s, "#") {
			newString = newString + sep + v
			sep = "\n"
			continue
		}
		newString = newString + sep + "# " + v
		sep = "\n"
	}
	return newString
}

func RemoveComments(s string) string {
	split := strings.Split(s, "\n")
	newString := ""
	sep := ""
	for _, v := range split {
		if strings.HasPrefix(s, "# ") {
			newString = newString + sep + v[2:]
			sep = "\n"
			continue
		}
		if strings.HasPrefix(s, "#") {
			newString = newString + sep + v[1:]
			sep = "\n"
			continue
		}
		newString = newString + sep + v
	}
	return newString
}
