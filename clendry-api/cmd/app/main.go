package main

import "github.com/artomsopun/clendry/clendry-api/internal/app"

const configsDir = "clendry-api/configs"

func main() {
	app.Run(configsDir)
}
