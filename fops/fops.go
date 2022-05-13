package fops

import (
	"github.com/cwxstat/dopt/samples"
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

	return newTag().addTagIfNeeded(filename)

}
