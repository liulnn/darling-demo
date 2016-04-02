package main

import (
	"darling"
	"darling-demo/plants"
)

func routerConfig(app *darling.App) {
	app.Handlers.Add("/v1/plants/(\\w+)", &plants.PlantCtrl{})
	app.Handlers.Add("/v1/plants", &plants.PlantsCtrl{})
}

func main() {
	var app = darling.NewApp()
	routerConfig(app)
	app.Run("127.0.0.1", 7432)
}
