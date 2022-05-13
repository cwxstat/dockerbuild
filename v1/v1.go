package v1

import (
	"github.com/cwxstat/dopt/fops"
)

func Init(file string) {
	t := fops.NewTag()
	t.AddTagIfNeeded(file)
}

func Update(file string) {
	t := fops.NewTag()
	t.Update(file)

}
