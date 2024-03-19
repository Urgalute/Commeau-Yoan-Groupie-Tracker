package controller

import (
	"BattlenetAPI/data"
	InitTemp "BattlenetAPI/temps"
	"fmt"
	"net/http"
	"strings"
)

// Affichage de la page collection avec toutes les spécialisations
func Collection(w http.ResponseWriter, r *http.Request) {

	route := r.URL.Query().Get("type")
	if route == "allclass" {
		data, err := GetAllClass()
		if err != nil {
			fmt.Println("Erreur lors de la récupération des classes : ", err)
			return
		}
		data.Route = route
		fmt.Println(data)
		InitTemp.Temp.ExecuteTemplate(w, "collection", data)
	}
	if route == "allrace" {
		data, err := GetAllRace()
		if err != nil {
			fmt.Println("Erreur lors de la récupération des races : ", err)
			return
		}
		data.Route = route
		fmt.Println(data)
		InitTemp.Temp.ExecuteTemplate(w, "collection", data)
	}
	

	if route == "tank" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecDetails []data.SpecDetails
		for _, spec := range TabSpec.TabSpec {
			data, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(data.SpecRole.RoleName), strings.ToLower("tank")) {
				TabSpecDetails = append(TabSpecDetails, data)

			}
		}
		fmt.Println(TabSpecDetails)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecDetails)
	}
	if route == "heal" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecDetails []data.SpecDetails
		for _, spec := range TabSpec.TabSpec {
			data, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(data.SpecRole.RoleName), strings.ToLower("soins")) {
				TabSpecDetails = append(TabSpecDetails, data)
				fmt.Println(TabSpecDetails)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecDetails)
	}
	if route == "dps" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecDetails []data.SpecDetails
		for _, spec := range TabSpec.TabSpec {
			data, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(data.SpecRole.RoleName), strings.ToLower("Dégâts")) {
				TabSpecDetails = append(TabSpecDetails, data)
				fmt.Println(TabSpecDetails)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecDetails)
	}
	if route == "force" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecDetails []data.SpecDetails
		for _, spec := range TabSpec.TabSpec {
			data, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(data.SpecPrimaryStat.PrimaryStatName), strings.ToLower("force")) {
				TabSpecDetails = append(TabSpecDetails, data)
				fmt.Println(TabSpecDetails)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecDetails)
	}
	if route == "agilite" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecDetails []data.SpecDetails
		for _, spec := range TabSpec.TabSpec {
			data, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(data.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Agilité")) {
				TabSpecDetails = append(TabSpecDetails, data)
				fmt.Println(TabSpecDetails)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecDetails)
	}
	if route == "intelligence" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecDetails []data.SpecDetails
		for _, spec := range TabSpec.TabSpec {
			data, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(data.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Intelligence")) {
				TabSpecDetails = append(TabSpecDetails, data)
				fmt.Println(TabSpecDetails)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecDetails)
	}
}
