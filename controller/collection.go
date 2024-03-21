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

	if route == "Tank" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}
		fmt.Println(route)

		var TabSpecData data.TabSpec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}
			TabSpec.Route = route
			if strings.Contains(strings.ToLower(dataSpec.SpecRole.RoleName), strings.ToLower("Tank")) {
				TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "Soins" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}
		fmt.Println(route)

		var TabSpecData data.TabSpec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}
			TabSpec.Route = route
			if strings.Contains(strings.ToLower(dataSpec.SpecRole.RoleName), strings.ToLower("Soins")) {
				TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "Dégâts" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}
		fmt.Println(route)

		var TabSpecData data.TabSpec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}
			TabSpec.Route = route
			if strings.Contains(strings.ToLower(dataSpec.SpecRole.RoleName), strings.ToLower("Dégâts")) {
				TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "Force" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}
		fmt.Println(route)

		var TabSpecData data.TabSpec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}
			TabSpec.Route = route
			if strings.Contains(strings.ToLower(dataSpec.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Force")) {
				TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "Agilité" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}
		fmt.Println(route)

		var TabSpecData data.TabSpec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}
			TabSpec.Route = route
			if strings.Contains(strings.ToLower(dataSpec.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Agilité")) {
				TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	}
	if route == "Intelligence" {
		TabSpec, listSpecErr := GetAllSpec()
		if listSpecErr != nil {
			// Erreur
		}
		fmt.Println(route)

		var TabSpecData data.TabSpec
		for _, spec := range TabSpec.TabSpec {
			dataSpec, errData := GetSpec(spec.Id)
			if errData != nil {
				fmt.Println("Erreur pour ID")
				continue
			}
			TabSpec.Route = route
			if strings.Contains(strings.ToLower(dataSpec.SpecPrimaryStat.PrimaryStatName), strings.ToLower("Intelligence")) {
				TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, route})
			}
		}
		fmt.Println(TabSpecData)
		InitTemp.Temp.ExecuteTemplate(w, "collection", TabSpecData)
	} else {
		//
	}
}
