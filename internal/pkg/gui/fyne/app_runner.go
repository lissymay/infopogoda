package fyne

import (
	"fyne.io/fyne/v2"
)

type appRunner struct {
	w fyne.Window
	a fyne.App
}

func NewAR(app fyne.App, w fyne.Window) *appRunner {
	return &appRunner{w: w, a: app}
}

func (ar *appRunner) Run() {
	ar.w.Show()
	ar.a.Run()
}
