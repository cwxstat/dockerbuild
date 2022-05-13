package yamlst

import (
	"fmt"
	"os"
	"strconv"
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
	DockerCMD []string    `yaml:"cmd"`
	Files     FilesUpdate `yaml:"changes,omitempty"`
}

type TopYaml struct {
	Image   string `yaml:"image"`
	Version string `yaml:"version"`
	Config  string `yaml:"local-config,omitempty"`
	Spec    Spec   `yaml:"spec,omitempty"`
}

func NewDY() *TopYaml {
	fu := FilesUpdate{
		Files: []string{"pod.yaml, dev-pod.yaml"},
	}
	t := &TopYaml{
		Image:   "us-central1-docker.pkg.dev/mchirico/public/septa",
		Version: "v0.0.5",
		Config:  "~/.docTag/config",
		Spec: Spec{
			Timestamp: time.Now(),
			Platform: Platform{
				OS: "linux/amd64",
			},
			Files: fu,
		},
	}
	return t
}

func (dy *TopYaml) Comments() (string, error) {
	a, err := yaml.Marshal(dy)
	if err != nil {
		return "", err
	}
	return addComments(string(a)), nil
}

func (dy *TopYaml) CommentsWithTags(tagBegin, tagEnd string) (string, error) {
	a, err := yaml.Marshal(dy)
	if err != nil {
		return "", err
	}
	s := tagBegin + "\n" + string(a) + "\n" + tagEnd + "\n"
	return addComments(s), nil
}

func (dy *TopYaml) UnMarshal(s string) error {
	ts := removeComments(s)
	err := yaml.Unmarshal([]byte(ts), dy)
	return err
}

func (dy *TopYaml) ImageVersion(image, version string) {
	dy.Image = image
	dy.Version = version

}

func (dy *TopYaml) NextMinor() error {

	if dy.Version == "" {
		dy.Version = "v0.0.1"
		return nil
	}
	// Assume v0.0.1
	split := strings.Split(dy.Version, ".")
	if len(split) != 3 {
		return fmt.Errorf("version not in `v0.0.1` format: %+v", dy.Version)
	}
	n, err := strconv.Atoi(split[2])
	if err != nil {
		return err
	}
	n = n + 1
	dy.Version = fmt.Sprintf("%s.%s.%d", split[0], split[1], n)
	return nil

}

func (dy *TopYaml) NextMajor() error {

	if dy.Version == "" {
		dy.Version = "v0.1.0"
		return nil
	}
	// Assume v0.0.1
	split := strings.Split(dy.Version, ".")
	if len(split) != 3 {
		return fmt.Errorf("version not in `v0.0.1` format: %+v", dy.Version)
	}
	n, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	n = n + 1
	dy.Version = fmt.Sprintf("%s.%d.%s", split[0], n, split[2])
	return nil

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

func addComments(s string) string {
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

func removeComments(s string) string {
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
