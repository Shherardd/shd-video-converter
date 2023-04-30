package main

import (
	"embed"
	video "shd-vc/Video"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	video_ := video.NewFileManager(&app.ctx)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "[shd] Video converter",
		Width:  900,
		Height: 580,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			video_,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
