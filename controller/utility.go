package controller

import (
	"BattlenetAPI/data"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var _ClientId string = "2bf60d3007414ebfa39fcf0137aa8d02"
var _SecretClient string = "RbOLgB8rALpg29Yy2d7mr1Qr1MCCP0ON"
var _Token string

type ResToken struct {
	Token string `json:"access_token"`
}

// Récupération du token d'autorisation
func AskToken() error {
	httpClient := http.Client{Timeout: 2 * time.Second}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, "https://oauth.battle.net/token?grant_type=client_credentials", nil)
	if errReq != nil {
		return fmt.Errorf("oupsss une erreur avec l'initialisation de la requete \"AskToken\"")
	}

	// Ajout des attributs dans l'en-tête de la requête
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", _ClientId, _SecretClient)))))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Envoie de la requête
	res, resErr := httpClient.Do(req)
	if resErr != nil {
		return fmt.Errorf("oups une erreur avec le résultat de la requete : \"AskToken\"")
	}

	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		dataRes, dataResErr := io.ReadAll(res.Body)
		if dataResErr != nil {
			return fmt.Errorf("oups une erreur durant la lecture des données de la réponse : \"AskToken\"")
		}

		var data ResToken
		errDecode := json.Unmarshal(dataRes, &data)
		if errDecode != nil {
			return fmt.Errorf("oups une erreur lors du décodage des données : \"AskToken\"")
		}

		_Token = fmt.Sprintf("Bearer %s", data.Token)
		fmt.Println(_Token)
		return nil
	}

	return fmt.Errorf("oupss une erreur code réponse %v : \"AskToken\"", res.StatusCode)
}

// Fonction qui récup toutes les spécialisation (Nom, Id)
func GetAllSpec() (data.TabSpec, error) {
	urlAPI := "https://eu.api.blizzard.com/data/wow/playable-specialization/index?namespace=static-eu&locale=fr_FR"

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return data.TabSpec{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return data.TabSpec{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		dataa, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return data.TabSpec{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes data.TabSpec
		errJson := json.Unmarshal(dataa, &DataRes)
		if errJson != nil {
			return data.TabSpec{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return data.TabSpec{}, fmt.Errorf("oupss une erreur avec la requete \" GetSpec \" code réponse : %v", res.StatusCode)
}

// Fonction qui récup toutes les classes (Nom, Id)
func GetAllClass() (data.TabClass, error) {
	urlAPI := "https://eu.api.blizzard.com/data/wow/playable-class/index?namespace=static-eu&locale=fr_FR"

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return data.TabClass{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return data.TabClass{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		dataa, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return data.TabClass{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes data.TabClass
		errJson := json.Unmarshal(dataa, &DataRes)
		if errJson != nil {
			return data.TabClass{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return data.TabClass{}, fmt.Errorf("oupss une erreur avec la requete \" GetClass \" code réponse : %v", res.StatusCode)
}

// Fonction qui récup toutes les races (Nom, Id)
func GetAllRace() (data.TabRace, error) {
	urlAPI := "https://eu.api.blizzard.com/data/wow/playable-race/index?namespace=static-eu&locale=fr_FR"

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Initialisation de la requête avec la méthode, l'endpoint et le corps de la requête
	req, _ := http.NewRequest(http.MethodGet, urlAPI, nil)

	// Ajout du token dans le header avec l'attribut 'Authorization'
	req.Header.Add("Authorization", _Token)

	// Envoie de la requête
	res, resErr := client.Do(req)
	if resErr != nil {
		return data.TabRace{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return data.TabRace{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		dataa, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return data.TabRace{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes data.TabRace
		errJson := json.Unmarshal(dataa, &DataRes)
		if errJson != nil {
			return data.TabRace{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return data.TabRace{}, fmt.Errorf("oupss une erreur avec la requete \" GetRace \" code réponse : %v", res.StatusCode)
}

func GetSpec() {

}
