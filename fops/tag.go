package fops

import (
	"errors"
	"os"

	"github.com/cwxstat/dopt/file"
	"github.com/cwxstat/dopt/uyaml"
)

var ErrReadTag = errors.New("read tag error")
var ErrWriteTag = errors.New("write tag error")

type tag struct {
	tagBeg string
	tagEnd string
	dy     *uyaml.TopYaml
}

func newTag() *tag {
	t := &tag{
		tagBeg: "<docb:",
		tagEnd: "</docb:",
		dy:     uyaml.NewDY(),
	}
	return t
}

func (t *tag) UpdateDY(dy *uyaml.TopYaml) *tag {
	t.dy = dy
	return t
}

func (t *tag)Dy() *uyaml.TopYaml {
	return t.dy
}

func (t *tag) addTagIfNeeded(filename string) error {
	s, err := file.Read(filename)
	if err != nil {
		return err
	}
	if _, _, err := file.GrabTag(s, t.tagBeg, t.tagEnd); err != nil {
		if err == file.ErrNoTag {
			if commentTag, err := t.dy.CommentsWithTags(t.tagBeg, t.tagEnd); err == nil {
				if f, err := file.HandleAppend(filename); err == nil {
					f.WriteString(commentTag)
					f.Close()
				}

			}
		}
		return nil
	}
	return err
}

func (t *tag) readTag(filename string) error {

	s, err := file.Read(filename)
	if err != nil {
		return err
	}

	if _, section, err := file.GrabTag(s, t.tagBeg, t.tagEnd); err == nil {
		err = t.dy.UnMarshal(section)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("read tag error")

}

func (t *tag) writeTag(filename string) error {
	s, err := file.Read(filename)
	if err != nil {
		return err
	}
	if dockerSection, _, err := file.GrabTag(s, t.tagBeg, t.tagEnd); err == nil {
		s := dockerSection + "\n"
		tag, err := t.dy.CommentsWithTags(t.tagBeg, t.tagEnd)
		if err != nil {
			return err
		}
		s += tag
		if err := os.WriteFile(filename, []byte(s), 0644); err != nil {
			return nil
		}
		return nil
	}
	return ErrWriteTag
}
