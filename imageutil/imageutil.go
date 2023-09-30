package imageutil

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"
)

type ImageData struct {
	Name   string
	Base64 string
}

func IsImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg"
}

func EncodeFileToBase64(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	base64Data := base64.StdEncoding.EncodeToString(data)
	return base64Data, nil
}
