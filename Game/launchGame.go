package game

import (
	"main/guistate"

	"fyne.io/fyne/v2/app"
)

func LaunchGame() {
	guistate.GameApp = app.New()
	guistate.GameWindow = guistate.GameApp.NewWindow("Blackjack")
	StartScreen()
	guistate.GameWindow.ShowAndRun()
}
