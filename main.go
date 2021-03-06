package main

import (
	"github.com/leaanthony/mewn"
	"github.com/pkg/browser"
	"github.com/wailsapp/wails"
	"log"
	"portfall/pkg/client"
)

// OpenInBrowser opens the operating system browser at the specified url
func OpenInBrowser(openUrl string) {
	err := browser.OpenURL(openUrl)
	if err != nil {
		log.Print(err)
	}
}

func main() {

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	c := &client.Client{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "Portfall",
		JS:        js,
		CSS:       css,
		Colour:    "#fff",
		Resizable: true,
	})
	app.Bind(c)
	app.Bind(OpenInBrowser)
	app.Run()
}
