package controlador

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func FolderContentsHandler(folderPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the hostname.
		hostName, err := os.Hostname()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Read the directory contents.
		files, err := ioutil.ReadDir(folderPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a slice of image data.
		var imageList []modelo.ImageData
		imageCount := 0

		// Iterate over the files and add image data to the slice.
		for _, file := range files {
			if isImageFile(file.Name()) {
				imagePath := filepath.Join(folderPath, file.Name())
				base64Data, err := encodeFileToBase64(imagePath)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				imageList = append(imageList, model.ImageData{
					Name:   file.Name(),
					Base64: base64Data,
				})

				imageCount++
				if imageCount >= 3 {
					break
				}
			}
		}

		// Create a data struct to pass to the template.
		data := struct {
			HostName   string
			FolderPath string
			Images     []model.ImageData
		}{
			HostName:   hostName,
			FolderPath: folderPath,
			Images:     imageList,
		}

		// Parse the template.
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the content type and execute the template.
		w.Header().Set("Content-Type", "text/html")
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
