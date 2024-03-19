package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func Categorie(w http.ResponseWriter, r *http.Request) {

	InitTemp.Temp.ExecuteTemplate(w, "categorie", nil)
}
