package main

import "github.com/artomsopun/clendry/clendry-api/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
