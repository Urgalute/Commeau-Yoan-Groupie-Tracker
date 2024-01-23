package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "index", nil)
}
