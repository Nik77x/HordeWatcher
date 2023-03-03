package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "KAIWorkerViz",
		Width:  1080,
		Height: 670,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 10},
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		Bind: []interface{}{
			app,
			//&Data.WorkerInfo{},
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
