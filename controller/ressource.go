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
	if route == "spec" {
		data, err := GetSpec(id)
		data.Route = route
		if err != nil {
			fmt.Println("Erreur lors de la récupération des spécialisations : ", err)
			return
		}
		InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
	}
	if route == "race" {
		data, err := GetRace(id)
		data.Route = route
		if err != nil {
			fmt.Println("Erreur lors de la récupération des races : ", err)
			return
		}
		InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
	}
	if route == "class" {
		data, err := GetClass(id)
		data.Route = route
		if err != nil {
			fmt.Println("Erreur lors de la récupération des classes : ", err)
			return
		}
		InitTemp.Temp.ExecuteTemplate(w, "ressource", data)
	}
}
