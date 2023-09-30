package imageutil

import (
	"encoding/base64"
	"io/ioutil"
)

// EncodeFileToBase64 toma una ruta de archivo y devuelve su contenido en formato base64.
func EncodeFileToBase64(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	base64Data := base64.StdEncoding.EncodeToString(data)
	return base64Data, nil
}
