package routeur

import (
	"BattlenetAPI/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {
	http.HandleFunc("/", controller.Index)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:3000/) - Server started on port:3000")
	http.ListenAndServe("localhost:3000", nil)
}