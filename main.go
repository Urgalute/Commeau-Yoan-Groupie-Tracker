package main

import (
	"BattlenetAPI/routeur"
	"BattlenetAPI/temps"
)

func main() {
	temps.InitTemplate()
	routeur.InitServ()
}
