package dockersdk

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

type Docker struct {
	cli        *client.Client
	tar        string
	image      string
	version    string
	platform   string
	dockerfile string
}

func NewDocker() (*Docker, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	docker := &Docker{
		cli:        cli,
		platform:   "linux/amd64",
		dockerfile: "Dockerfile",
	}
	return docker, err

}

// dopt/
func (d *Docker) Tar(dir string) *Docker {
	d.tar = dir
	return d
}

func (d *Docker) Image(image string) *Docker {
	d.image = image
	return d
}

func (d *Docker) Version(version string) *Docker {
	d.version = version
	return d
}

func (d *Docker) Platform(platform string) *Docker {
	d.platform = platform
	return d
}

func (d *Docker) Dockerfile(file string) *Docker {
	d.dockerfile = file
	return d
}

func (d *Docker) ImageBuild() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	tar, err := archive.TarWithOptions(d.tar, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: d.dockerfile,
		Tags:       []string{d.image + d.version},
		NoCache:    true,
		Platform:   d.platform,
		Remove:     true,
	}
	res, err := d.cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = print(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func print(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
