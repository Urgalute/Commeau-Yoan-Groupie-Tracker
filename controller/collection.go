package controller

import (
	InitTemp "BattlenetAPI/temps"
	"fmt"
	"net/http"
)

// Affichage de la page collection avec toutes les spécialisations
func Collection(w http.ResponseWriter, r *http.Request) {
	data, err := GetAllSpec()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des spécialisations : ", err)
		return
	}
	InitTemp.Temp.ExecuteTemplate(w, "collection", data)
}
