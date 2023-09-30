package main

import (
	"flag"
	"fmt"
	"net/http"
	//"github.com/yourusername/myproject/web"
)

func main() {
	folderPathPtr := flag.String("folder", "", "Ruta de la carpeta")
	portPtr := flag.Int("port", 8081, "Número de puerto para el servidor HTTP")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		web.FolderContentsHandler(w, r, *folderPathPtr)
	})

	addr := fmt.Sprintf(":%d", *portPtr)
	fmt.Printf("Servidor HTTP en ejecución en http://localhost%s\n", addr)
	http.ListenAndServe(addr, nil)
}
