package main

import (
	"flag"
	"fmt"
	"net/http"

	//"github.com/bard/golang/image-server/handler"
	//"github.com/bard/golang/image-server/model"
	//"github.com"

	"github.com/SSabogalG/Computacion-nube/handler"
)

func main() {
	// Define the command line flags.
	folderPathPtr := flag.String("folder", "", "Ruta de la carpeta")
	portPtr := flag.Int("port", 8081, "Número de puerto para el servidor HTTP")
	flag.Parse()

	// Register the folder contents handler.
	http.HandleFunc("/", handler.FolderContentsHandler(folderPathPtr))

	// Start the HTTP server.
	addr := fmt.Sprintf(":%d", *portPtr)
	fmt.Printf("Servidor HTTP en ejecución en http://localhost%s\n", addr)
	http.ListenAndServe(addr, nil)
}
