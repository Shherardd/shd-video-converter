package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet() {
	openDialogOptions := runtime.OpenDialogOptions{
		Title:            "Choose a directory",
		DefaultDirectory: "/Users/sherard/Dev",
	}

	path, err := runtime.OpenDirectoryDialog(a.ctx, openDialogOptions)
	if err != nil {
		fmt.Errorf("error opening directory dialog: %s", err)
	}
	fmt.Println(path)
}
