package main

import (
	"embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/riadafridishibly/vault-app/internal/imghandler"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = wails.Run(&options.App{
		Title:    "Vault",
		Width:    1024,
		Height:   768,
		MinWidth: 1024,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: imghandler.NewImageHandler(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app.dialogService,
			app.passwordService,
			app.databaseService,
			app.clipboardService,
			app.assetServer,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
