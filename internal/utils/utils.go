package utils

import (
	"io/ioutil"
)

func ReadFileContent(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
