package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "about", nil)
}
