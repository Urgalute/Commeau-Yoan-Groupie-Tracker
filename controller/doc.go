package controller

import (
	InitTemp "BattlenetAPI/temps"
	"net/http"
)

func Doc(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "doc", nil)
}
