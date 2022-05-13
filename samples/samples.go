package samples

import (
	"errors"
	
	img "github.com/cwxstat/dockerbuild/samples/dockerimages"
	"os"
)
var ErrFileExists = errors.New("sample file exists")
func CreateSample(file string) error {
	if _, err := os.Stat(file); err == nil {
		// file exist
		return ErrFileExists

	} else if errors.Is(err, os.ErrNotExist) {

		d1 := []byte(img.Images("golang"))
		err = os.WriteFile(file, d1, 0644)
		if err != nil {
			return err
		}

	}
	return nil
}
