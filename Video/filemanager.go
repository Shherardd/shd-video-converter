package video

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"time"
)

type FileManager struct {
	ctx *context.Context

	File           File
	OuputDirectory string
}

func NewFileManager(ctx *context.Context) *FileManager {
	return &FileManager{
		ctx: ctx,
	}
}

func (v *FileManager) GetHenlo() string {
	v.SendProgressFake()
	return "henlo!"
}

func (v *FileManager) GetFile() *File {
	return &v.File
}

func (v *FileManager) ChooseFile() string {
	var defaultDir string
	if v.File.Path != "" {
		defaultDir = v.File.Path
	} else {
		defaultDir, _ = os.UserHomeDir()
	}

	openDialogOptions := runtime.OpenDialogOptions{
		Title:            "Choose a file",
		DefaultDirectory: defaultDir,
		Filters: []runtime.FileFilter{
			{"Videos", ".mov;*.mp4"},
		},
	}

	path, err := runtime.OpenFileDialog(*v.ctx, openDialogOptions)
	if err != nil {
		fmt.Errorf("error opening file dialog: %s", err)
	}
	v.File.Path = path
	return path
}

func (v *FileManager) ChooseDirectory() string {
	openDialogOptions := runtime.OpenDialogOptions{
		Title:            "Choose a directory",
		DefaultDirectory: "/Users/sherard/Dev",
	}

	dir, err := runtime.OpenDirectoryDialog(*v.ctx, openDialogOptions)
	if err != nil {
		fmt.Errorf("error opening directory dialog: %s", err)
	}
	v.OuputDirectory = dir
	return dir
}

func (v *FileManager) SendProgress(progress float64) {
	// Enviar el progreso al proceso renderer
	runtime.EventsEmit(*v.ctx, "progress", progress)

	//runtime.Window().Send("progress", progress)
}

// a SendProgress implementation that send progress from 0 to 100 in 10 steps (10, 20, 30, ..., 100)
func (v *FileManager) SendProgressFake() {
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		v.SendProgress(float64(i * 10))
	}
}
