package main

import (
	"BattlenetAPI/controller"
	"BattlenetAPI/routeur"
	"BattlenetAPI/temps"
)

func main() {
	controller.FirstRequest()
	temps.InitTemplate()
	routeur.InitServ()
}
