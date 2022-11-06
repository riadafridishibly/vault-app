package main

import (
	"embed"

	"github.com/riadafridishibly/vault-app/pkg/imghandler"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "Vault",
		Width:            1024,
		Height:           768,
		MinWidth:         1024,
		Assets:           assets,
		AssetsHandler:    imghandler.NewImageHandler(),
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
