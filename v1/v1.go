package v1

import (
	"github.com/cwxstat/dopt/tag"
)

func Init(file string) {
	t := tag.NewTag()
	t.AddTagIfNeeded(file)
}

func Update(file string) {
	t := tag.NewTag()
	t.Update(file)
}
