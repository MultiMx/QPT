package util

import (
	"io"
	"os"
)

func SaveTmpPack(i io.ReadCloser) (string, error) {
	defer i.Close()

	file, e := os.CreateTemp("", "qpt-install-*****.tar.gz")
	if e != nil {
		return "", e
	}
	defer file.Close()

	_, e = io.Copy(file, i)
	return file.Name(), e
}
