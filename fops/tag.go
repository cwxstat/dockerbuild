package fops

import (
	"github.com/cwxstat/dopt/file"
	"github.com/cwxstat/dopt/uyaml"
)

type tag struct {
	tagBeg string
	tagEnd string
}

func newTag() *tag {
	t := &tag{
		tagBeg: "<docb:",
		tagEnd: "</docb:",
	}
	return t
}

func (t *tag) addTagIfNeeded(filename string) error {
	s, err := file.Read(filename)
	if err != nil {
		return err
	}
	if _, _, err := file.GrabTab(s, t.tagBeg, t.tagEnd); err != nil {
		if err == file.ErrNoTag {
			dy := uyaml.NewDY()
			if commentTag, err := dy.Comments(); err == nil {
				if f, err := file.HandleAppend(filename); err == nil {
					s := "# " + t.tagBeg + ">\n" + commentTag + "\n# " + t.tagEnd + ">"
					f.WriteString(s)
					f.Close()
				}

			}
		}
		return nil
	}
	return err
}
