package fops

import (
	"github.com/cwxstat/dutil/file"
	"github.com/cwxstat/dutil/samples"
	"github.com/cwxstat/dutil/uyaml"
)

type FOPS struct {
}

func NewFOPS() *FOPS {
	return &FOPS{}
}

func addTagIfNeeded(filename string) error {
	s, err := file.Read(filename)
	if err != nil {
		return err
	}
	if _, _, err := file.GrabTab(s, "<docb:", "</docb:"); err != nil {
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
		return nil
	}
	return err
}

func (f *FOPS) Sample() error {

	filename := "Dockerfile.sample"
	if err := samples.CreateSample(filename); err != nil {
		if err != samples.ErrFileExists {
			return err
		}
	}

	return addTagIfNeeded(filename)

}
