package main

import (
	"battery-notification/internal/settings"
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	mySettings, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "Battery Notification",
		JS:        js,
		CSS:       css,
		Resizable: true,
		Colour:    "#131313",
	})
	app.Bind(mySettings)
	app.Run()
}
