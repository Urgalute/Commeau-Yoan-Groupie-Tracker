package routeur

import (
	"BattlenetAPI/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {
	http.HandleFunc("/", controller.Index)
	//http.HandleFunc("/doc", )
	//http.HandleFunc("/about", )
	http.HandleFunc("/collection", controller.Collection)
	http.HandleFunc("/categorie", controller.Categorie)
	//http.HandleFunc("/favoris", )
	http.HandleFunc("/ressource", controller.Ressource)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:3000/) - Server started on port:3000")
	http.ListenAndServe("localhost:3000", nil)
}
