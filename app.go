package main

import (
	"darling"
)

var app = darling.NewApp()

func main() {
	app.Run("127.0.0.1", 7432)
}
