package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	video "shd-vc/Video"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const debugMode = false

func main() {
	app := NewApp()
	video_ := video.NewFileManager(&app.ctx)

	debugOptions := options.Debug{
		OpenInspectorOnStartup: debugMode,
	}

	err := wails.Run(&options.App{
		Title:       "shd Video converter",
		Width:       900,
		Height:      580,
		StartHidden: false,
		Debug:       debugOptions,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		OnStartup:        app.startup,
		Mac: &mac.Options{
			WebviewIsTransparent: true,
		},
		Bind: []interface{}{
			app,
			video_,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
