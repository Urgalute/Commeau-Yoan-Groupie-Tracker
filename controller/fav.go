package controller

import (
	"BattlenetAPI/data"
	InitTemp "BattlenetAPI/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var FavList []data.Fav

func Favoris(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	data := FavList
	fmt.Println(FavList)
	InitTemp.Temp.ExecuteTemplate(w, "fav", data)
}

func FavTreatment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("type") == "add" {
		route := r.URL.Query().Get("route")
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			//err
		}
		name := r.URL.Query().Get("name")
		var favoris data.Fav
		if route == "spec" && IdExist(id) {
			favoris.Spec.Name = name
			favoris.Spec.Id = id
			favoris.FavId = GetFavId()
			favoris.Route = route
			favoris.Spec.Route = route
			AddFavoris(favoris, true)
		}
		http.Redirect(w, r, "/favoris", http.StatusSeeOther)
	} else if r.URL.Query().Get("type") == "del" {
		FavId, err := strconv.Atoi(r.URL.Query().Get("favid"))
		if err != nil {
			//err
		}
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			//err
		}
		if !IdExist(int(id)) {
			RemoveFav(int(FavId), true)
		}
		http.Redirect(w, r, "/favoris", http.StatusSeeOther)
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

// Verifie si la spec est déjà présente dans les Fav
func IdExist(id int) bool {
	for _, Fav := range FavList {
		if Fav.Spec.Id == id {
			return false
		}
	}
	return true
}
func GetFavId() int {
	id := len(FavList) + 1
	return id
}

func RemoveFav(index int, save bool) {
	GetDataFromJson()
	var NewFav []data.Fav
	for _, fav := range FavList {
		if fav.FavId != index {
			NewFav = append(NewFav, fav)
		}
	}
	FavList = NewFav
	if save {
		SetDataToJson()
	}
}
func GetAllFav() []data.Fav {
	GetDataFromJson()
	return FavList
}
