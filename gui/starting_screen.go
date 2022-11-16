package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func StartScreen() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Blackjack")
	image := canvas.NewImageFromFile("gui/blacklogo-removebg-preview.png")
	image.FillMode = canvas.ImageFillOriginal
	image.Show()

	//welcome := canvas.NewText("Welcome to Blackjack!", color.White)
	centeredImage := container.NewCenter(image)
	startButton := widget.NewButton("Start Game", func() {
		//difficultyWindow := myApp.NewWindow("Difficulty Selection")
		log.Println("start button tapped")

		difficultySelector := widget.NewLabel("Select the aggressiveness of the AI Dealer: ")

		radio := widget.NewRadioGroup([]string{"Mild", "Moderate", "Aggressive"}, func(value string) {
			log.Println("Radio set to", value)
			myWindow.SetContent(GameScreen())
		})

		myWindow.SetContent(container.NewVBox(difficultySelector, radio))
	})

	quit_button := widget.NewButton("Quit", func() {
		log.Println("tapped")
	})

	game_buttons := container.New(layout.NewHBoxLayout(), startButton, quit_button)
	bottom := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), game_buttons, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), centeredImage, layout.NewSpacer(), bottom))
	myWindow.ShowAndRun()
}
