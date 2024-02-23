package main

import (
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

// Cette méthode permet de récupérer un Token (il faudra initialiser les variables _ClientId & _SecretClient)
// Elle retourne une "erreur" lorsqu'une erreur se produit lors de la requête sinon elle retourne "nil"
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

type Char struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type CallAPI struct {
	ListChar []Char `json:"character_specializations"`
}

func GetPlayer() (CallAPI, error) {
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
		return CallAPI{}, fmt.Errorf("oupss erreur avec la requete")
	}

	if res.StatusCode == http.StatusUnauthorized {
		return CallAPI{}, fmt.Errorf("oupss erreur d'authentification : code 401")
		// Ici, il faudra gérer le token expiré...
	} else if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		data, dataErr := io.ReadAll(res.Body)
		if dataErr != nil {
			fmt.Println("err lecture body")
			return CallAPI{}, fmt.Errorf("oupss une erreur lors de lecture de la réponse")
		}

		var DataRes CallAPI
		errJson := json.Unmarshal(data, &DataRes)
		if errJson != nil {
			return CallAPI{}, fmt.Errorf("oupss une erreur lors du decodage des données")
		}
		return DataRes, nil
	}

	return CallAPI{}, fmt.Errorf("oupss une erreur avec la requete \" GetPlayer \" code réponse : %v", res.StatusCode)
}

func main1() {
	errToken := AskToken()
	if errToken != nil {
		fmt.Println(errToken.Error())
		return
	}

	data, errData := GetPlayer()
	if errData != nil {
		fmt.Println(errData.Error())
		return
	}

	fmt.Println(data)
}
