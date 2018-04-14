package utils

import (
	"testing"
	"os"
)

func TestOptions_ConvertImage(t *testing.T) {
		option := Options{"./", "./", "jpeg", "png"}

		dir, _ := os.Getwd()

		contents := dir + "/hoge.png"
		_, err := os.Stat(contents)

		option.ConvertImage()

		if !os.IsExist(err) {
			t.Fatalf(`convert png file not exist but it is %s`, contents)
		}
}

