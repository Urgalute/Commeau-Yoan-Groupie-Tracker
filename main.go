package main

import (
	"BattlenetAPI/controller"
	"BattlenetAPI/routeur"
	"BattlenetAPI/temps"
)

func main() {
	controller.AskToken()
	temps.InitTemplate()
	routeur.InitServ()
}
