package tag

import (
	"errors"
	"os"

	"github.com/cwxstat/dopt/constants"
	"github.com/cwxstat/dopt/file"
	"github.com/cwxstat/dopt/yamlst"
)

var ErrReadTag = errors.New("read tag error")
var ErrWriteTag = errors.New("write tag error")

type tag struct {
	tagBeg   string
	tagEnd   string
	filename string
	//dy       *yamlst.TopYaml
	dy yaml
}

func NewTag() *tag {
	t := &tag{
		tagBeg: constants.TAG_BEGIN,
		tagEnd: constants.TAG_END,
		dy:     yamlst.NewDY(),
	}
	return t
}

func newTag() *tag {
	t := &tag{
		tagBeg: constants.TAG_BEGIN,
		tagEnd: constants.TAG_END,
		dy:     yamlst.NewDY(),
	}
	return t
}

func (t *tag) UpdateDY(dy *yamlst.TopYaml) *tag {
	t.dy = dy
	return t
}



func (t *tag) AddTagIfNeeded(filename string) {
	t.addTagIfNeeded(filename)
	t.filename = filename
}

func (t *tag) addTagIfNeeded(filename string) error {
	t.filename = filename
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

func (t *tag) Update(filename string) {
	t.filename = filename
	t.readTag(filename)
	t.dy.NextMinor()
	t.writeTag(filename)

}

func (t *tag) Read(filename string) error {
	t.filename = filename
	return t.readTag(filename)
}

func (t *tag) ImageVersion(image, version string) error {
	t.dy.ImageVersion(image, version)
	t.writeTag(t.filename)
	return nil

}

func (t *tag) readTag(filename string) error {
	t.filename = filename
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
	t.filename = filename
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
