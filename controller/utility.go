package controller

import (
	"BattlenetAPI/data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*const (
	clientID     = "2bf60d3007414ebfa39fcf0137aa8d02"
	clientSecret = "RbOLgB8rALpg29Yy2d7mr1Qr1MCCP0ON"
	redirectURI  = "http://localhost:3000/"
)*/

func FirstRequest() {
	Token := "Bearer EURjuBhTPeiTV75N4EOp906wotOJmgW4wQ"
	urlApi := "https://eu.api.blizzard.com/data/wow/playable-specialization/index?namespace=static-eu&locale=fr_FR"
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oups il y a un problème avec la requête : ", errReq.Error())
	}
	req.Header.Set("Authorization", Token)

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("Oups il y a un problème avec l'envoie de la requête : ", errRes.Error())
	}
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oups il y a un problème avec la lecture du corps : ", errBody.Error())
	}
	var decodeData data.TabSpec
	json.Unmarshal(body, &decodeData)
	fmt.Println(decodeData)
}
