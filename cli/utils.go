package cli

import (
	"io/ioutil"
)

func readFile(path string) ([]byte, error) {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		return []byte(""), err
	}

	return contents, nil
}

func getAbsolutePath(path string) (string, error) {
	absolutePath, err := filepathAbs(path)

	if err != nil {
		return "", err
	}

	return absolutePath, nil
}
