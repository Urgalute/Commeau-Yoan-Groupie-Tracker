package controller

import (
	"BattlenetAPI/data"
	InitTemp "BattlenetAPI/temps"
	"fmt"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_ = r.ParseForm()
		data := SearchName(r.FormValue("query"))
		InitTemp.Temp.ExecuteTemplate(w, "recherche", data)
	}
}

func SearchName(name string) data.TabSpec {
	TabSpec, listSpecErr := GetAllSpec()
	if listSpecErr != nil {
		// Erreur
	}

	var TabSpecData data.TabSpec
	for _, spec := range TabSpec.TabSpec {
		dataSpec, errData := GetSpec(spec.Id)
		if errData != nil {
			fmt.Println("Erreur pour ID")
			continue
		}
		if strings.Contains(strings.ToLower(dataSpec.Name), strings.ToLower(name)) {
			TabSpecData.TabSpec = append(TabSpecData.TabSpec, data.Spec{dataSpec.Name, dataSpec.Id, ""})
		}
	}
	fmt.Println(TabSpecData)
	return TabSpecData
}

