package controller

import (
	InitTemp "BattlenetAPI/temps"
	"fmt"
	"net/http"
)

func Categorie(w http.ResponseWriter, r *http.Request) {
	data, err := GetAllClass()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des classes : ", err)
		return
	}
	InitTemp.Temp.ExecuteTemplate(w, "categorie", data)
}
