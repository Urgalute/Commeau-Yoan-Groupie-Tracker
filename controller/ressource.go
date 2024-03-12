package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func Ressource(w http.ResponseWriter, r *http.Request) {
	//Data := 
	InitTemp.Temp.ExecuteTemplate(w, "ressource", nil)
}
