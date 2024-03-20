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
		fmt.Println(route)

		var TabSpecData []data.Spec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(dataSpec.SpecRole.RoleName), strings.ToLower("tank")) {
				TabSpecData = append(TabSpecData, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "heal" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecData []data.Spec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(dataSpec.SpecRole.RoleName), strings.ToLower("soins")) {
				TabSpecData = append(TabSpecData, data.Spec{dataSpec.Name, dataSpec.Id, route})
				fmt.Println(TabSpecData)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "dps" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecData []data.Spec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(dataSpec.SpecRole.RoleName), strings.ToLower("Dégâts")) {
				TabSpecData = append(TabSpecData, data.Spec{dataSpec.Name, dataSpec.Id, route})
				fmt.Println(TabSpecData)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "force" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecData []data.Spec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(dataSpec.SpecPrimaryStat.PrimaryStatName), strings.ToLower("force")) {
				TabSpecData = append(TabSpecData, data.Spec{dataSpec.Name, dataSpec.Id, route})
				fmt.Println(TabSpecData)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "agilite" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecData []data.Spec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(dataSpec.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Agilité")) {
				TabSpecData = append(TabSpecData, data.Spec{dataSpec.Name, dataSpec.Id, route})
				fmt.Println(TabSpecData)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "intelligence" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}

		var TabSpecData []data.Spec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}

			if strings.Contains(strings.ToLower(dataSpec.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Intelligence")) {
				TabSpecData = append(TabSpecData, data.Spec{dataSpec.Name, dataSpec.Id, route})
				fmt.Println(TabSpecData)
			}
		}

		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
}
