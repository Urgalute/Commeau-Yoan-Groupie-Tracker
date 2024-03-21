package controller

import (
	"BattlenetAPI/data"
	InitTemp "BattlenetAPI/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)
var FavList []data.Fav

func Fav(w http.ResponseWriter, r *http.Request) {
		GetDataFromJson()
		data := FavList
	InitTemp.Temp.ExecuteTemplate(w, "fav", data)
}




func AddFav() {
	var route string
	var Favoris data.Fav
	if route == "spec" {
		Favoris.Id = 123
		Favoris.FavId = GetFavId()
		Favoris.Route = ""
		AddFavoris(Favoris, true)
	}
}
func AddFavoris(Favoris data.Fav, save bool) {
	GetDataFromJson()
	FavList = append(FavList, Favoris)
	if save {
		SetDataToJson()
	}
}

func GetDataFromJson() {
	data, err := os.ReadFile("data/data.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &FavList) //passage en json vers la struct
}

func SetDataToJson() {
	data, err := json.Marshal(FavList)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/data.json", data, 0644)
}

func GetFavId() int {
	id := len(FavList) + 1
	return id
}