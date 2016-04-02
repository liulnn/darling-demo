package main

import (
	"darling"
	"darling-demo/plants"
)

func main() {
	var app = darling.NewApp()
	app.Handlers.Add("/v1/plants", &plants.PlantsCtrl{})
	app.Run("127.0.0.1", 7432)
}
