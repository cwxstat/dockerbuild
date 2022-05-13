package fops

import (
	"fmt"
	"github.com/cwxstat/dockerbuild/file"
	"github.com/cwxstat/dockerbuild/samples"
	"github.com/cwxstat/dockerbuild/uyaml"
)

type FOPS struct {
}

func NewFOPS() *FOPS {
	return &FOPS{}
}

func (f *FOPS) Sample() error {

	filename := "Dockerfile.sample"
	if err := samples.CreateSample(filename); err != nil {
		if err != samples.ErrFileExists {
			return err
		}
	}

	s, err := file.Read(filename)
	if err != nil {
		return err
	}
	if r, tag, err := file.GrabTab(s, "<docb:", "</docb:"); err != nil {
		if err == file.ErrNoTag {
			dy := uyaml.NewDY()
			if commentTag, err := dy.Comments(); err == nil {
				if f, err := file.HandleAppend(filename); err == nil {
					s := "# <docb:>\n" + commentTag + "\n# </docb:>"
					f.WriteString(s)
					f.Close()
				}

			}
		}
		fmt.Println(r, tag)
	} else {
		return err
	}
	return nil
}
