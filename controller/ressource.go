package controller

import (
	InitTemp "BattlenetAPI/temps"
	"fmt"
	"net/http"
	"strconv"
)

func Ressource(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Query().Get("route")
	id, err := strconv.Atoi(r.URL.Query().Get("type"))
	fmt.Println(route)
	fmt.Println(id)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	if route == "allclass" {
		data, err := GetClass(id)
		fmt.Println(id)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des classes : ", err)
			return
		}
		InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
	}
	if route == "allrace" {
		data, err := GetRace(id)
		fmt.Println(id)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des races : ", err)
			return
		}
		InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
	}
}

/*func ClassRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data, err := GetClass(id)
	fmt.Println(id)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des classes : ", err)
		return
	}
	InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
}

func RaceRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("type"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data, err := GetRace(id)
	fmt.Println(id)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des races : ", err)
		return
	}
	InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
}
*/
