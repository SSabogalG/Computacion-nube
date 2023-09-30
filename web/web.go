package web

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func FolderContentsHandler(w http.ResponseWriter, r *http.Request, folderPath string) {
	hostName, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var imageList []ImageData
	imageCount := 0 // Contador de imágenes

	for _, file := range files {
		if IsImageFile(file.Name()) {
			imagePath := filepath.Join(folderPath, file.Name())
			base64Data, err := EncodeFileToBase64(imagePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			imageList = append(imageList, ImageData{
				Name:   file.Name(),
				Base64: base64Data,
			})

			imageCount++
			if imageCount >= 3 {
				break // Limitar a 3 imágenes
			}
		}
	}

	data := struct {
		HostName   string
		FolderPath string
		Images     []ImageData
	}{
		HostName:   hostName,
		FolderPath: folderPath,
		Images:     imageList,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
