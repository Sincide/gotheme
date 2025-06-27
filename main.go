package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"net/http"
	"os/exec"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	a := app.New()
	w := a.NewWindow("GoTheme")

	selectWallpaper := func() {
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil || uc == nil {
				return
			}
			path := uc.URI().Path()
			applyWallpaper(path)
			palette := extractPalette(path)
			showPalette(w, palette)
			go queryAI(palette)
		}, w)
	}

	selectButton := fyne.NewMenuItem("Select Wallpaper", func() { selectWallpaper() })
	menu := fyne.NewMainMenu(fyne.NewMenu("File", selectButton))
	w.SetMainMenu(menu)

	w.SetContent(container.NewVBox())
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func applyWallpaper(path string) {
	exec.Command("swww", "img", path).Run()
}

func extractPalette(path string) []color.Color {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		return nil
	}
	colors := colorful.FastKMeans(img, 8)
	var palette []color.Color
	for _, c := range colors {
		palette = append(palette, c)
	}
	return palette
}

func showPalette(w fyne.Window, palette []color.Color) {
	if palette == nil {
		return
	}
	var objects []fyne.CanvasObject
	for _, c := range palette {
		rect := canvas.NewRectangle(c)
		rect.SetMinSize(fyne.NewSize(40, 40))
		objects = append(objects, rect)
	}
	w.SetContent(container.NewHBox(objects...))
}

func queryAI(palette []color.Color) {
	var hex []string
	for _, c := range palette {
		hex = append(hex, colorful.MakeColor(c).Hex())
	}
	p := fmt.Sprintf("Analyze palette: %v", hex)
	jsonData := fmt.Sprintf(`{"model":"llama2","prompt":"%s","stream":false}`, p)
	http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer([]byte(jsonData)))
}
