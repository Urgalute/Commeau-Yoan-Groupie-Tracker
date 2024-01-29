package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func Collection(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "colleection", nil)
}