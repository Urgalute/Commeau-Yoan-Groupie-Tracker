package controller

import (
	"fmt"
	"net/http"
)

const (
	clientID     = "2bf60d3007414ebfa39fcf0137aa8d02"
	clientSecret = "RbOLgB8rALpg29Yy2d7mr1Qr1MCCP0ON"
	redirectURI  = "http://localhost:3000/"
)

func RequestToken(w http.ResponseWriter, r *http.Request) {

	authURL := fmt.Sprintf("https://us.battle.net/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", clientID, redirectURI)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)

}
