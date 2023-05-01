package video

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
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

func (v *FileManager) ChooseFile() (string, error) {
	fmt.Println("Choose file from back")
	var defaultDir string
	defaultDir, _ = os.UserHomeDir()

	openDialogOptions := runtime.OpenDialogOptions{
		Title:            "Choose a file",
		DefaultDirectory: defaultDir + "/Movies",
	}

	path, err := runtime.OpenFileDialog(*v.ctx, openDialogOptions)
	if err != nil {
		fmt.Println("error opening file dialog: ", err)
		return "", err
	}

	v.File.Path = path

	return path, nil
}

func (v *FileManager) ChooseDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error al obtener el directorio home:", err)
		return "Error al obtener el directorio home"
	}

	openDialogOptions := runtime.OpenDialogOptions{
		Title:            "Choose a directory",
		DefaultDirectory: home + "/Movies",
	}

	dir, err := runtime.OpenDirectoryDialog(*v.ctx, openDialogOptions)
	if err != nil {
		fmt.Errorf("error opening directory dialog: %s", err)
	}
	v.OuputDirectory = dir
	return dir
}

func (v *FileManager) Convert() (string, error) {
	movPath := v.File.Path

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	randNum := r.Int()
	mp4Path := v.OuputDirectory + "/out_" + strconv.Itoa(randNum) + ".mp4"

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error al obtener el directorio home:", err)
		return "error", err
	}

	ffmpegPath := home + "/shd/lib/ffmpeg"

	cmd := exec.Command(ffmpegPath, "-i", movPath, mp4Path)

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al convertir el archivo:", err)
		return "error", err
	}

	fmt.Println("Archivo convertido con éxito!")
	return "success", nil
}

func (v *FileManager) GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error al obtener el directorio home:", err)
		return "error", err
	}
	v.OuputDirectory = homeDir
	return homeDir, nil
}

func (v *FileManager) SendProgress(progress float64) {
	runtime.EventsEmit(*v.ctx, "progress", progress)
}

func (v *FileManager) SendProgressFake() {
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		v.SendProgress(float64(i * 10))
	}
}
