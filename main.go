package main

import (
	"github.com/Reljod/Reljod-Portfolio-Backend/app"
	"github.com/Reljod/Reljod-Portfolio-Backend/swagger"
)

func main() {
	swagger.Setup()
	app.SetupRouter()
}
