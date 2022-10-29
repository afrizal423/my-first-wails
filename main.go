package main

import (
	"changeme/spamin"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	spam := spamin.GasSpamin()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Nyepamin Secreto",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        spam.Startup,
		Bind: []interface{}{
			spam,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
