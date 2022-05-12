package samples

import (
	"errors"
	"fmt"
	img "github.com/cwxstat/dockerbuild/samples/dockerimages"
	"os"
)

func CreateSample(file string) error {
	if _, err := os.Stat(file); err == nil {
		// file exist
		return fmt.Errorf("File exists: %s", file)

	} else if errors.Is(err, os.ErrNotExist) {

		d1 := []byte(img.Images("golang"))
		err = os.WriteFile(file, d1, 0644)
		if err != nil {
			return err
		}

	}
	return nil
}
