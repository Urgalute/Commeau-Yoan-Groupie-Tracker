package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func Fav(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "fav", nil)
}