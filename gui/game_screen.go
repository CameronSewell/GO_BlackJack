package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GameScreen() *fyne.Container {

	//Set of buttons and container for them
	hitButton := widget.NewButton("Hit", func() { player.hit() })
	splitButton := widget.NewButton("Split", func() { player.split() })
	doubleButton := widget.NewButton("Double", func() { player.double() })
	surrenderButton := widget.NewButton("Surrender", func(){ player.quit() })
	standButton := widget.NewButton("Stand", func() { player.stand() })
	gameButtons := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), layout.NewSpacer(), hitButton, surrenderButton, doubleButton, splitButton,
		standButton, layout.NewSpacer(), layout.NewSpacer())

	//Game's labels
	betLabel := widget.NewLabel("Bet: ")
	totalPotLabel := widget.NewLabel("$: ")
	handLabel := widget.NewLabel("Your hand: ")
	dealerHandLabel := widget.NewLabel("Dealer's hand: ")

	//Aligning game labels left
	betLabelAlignment := container.NewHBox(betLabel, layout.NewSpacer(), layout.NewSpacer())
	totalPotLabelAlignment := container.NewHBox(totalPotLabel, layout.NewSpacer(), layout.NewSpacer())
	handLabelAlignment := container.NewHBox(layout.NewSpacer(), handLabel, layout.NewSpacer())
	dealerLabelAlignment := container.NewHBox(layout.NewSpacer(), dealerHandLabel, layout.NewSpacer())

	//Player card set up
	firstCardImage := chooseCard("5", "diamonds")
	secondCardImage := chooseCard("king", "clubs")
	firstCardImage.FillMode = canvas.ImageFillOriginal
	secondCardImage.FillMode = canvas.ImageFillOriginal
	cardAlignment := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), firstCardImage, secondCardImage, layout.NewSpacer())

	//Dealer card set up
	thirdCardImage := chooseCard("2", "hearts")
	fourthCardImage := chooseCard("queen", "hearts")
	thirdCardImage.FillMode = canvas.ImageFillOriginal
	fourthCardImage.FillMode = canvas.ImageFillOriginal
	dealerCardAlignment := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), firstCardImage, secondCardImage, layout.NewSpacer())

	//combining labels with cards
	labelAlignment := container.New(layout.NewVBoxLayout(), betLabelAlignment, totalPotLabelAlignment,
		dealerLabelAlignment, dealerCardAlignment, handLabelAlignment, cardAlignment)

	//Putting it all together
	retVal := container.New(layout.NewVBoxLayout(), labelAlignment, layout.NewSpacer(), layout.NewSpacer(), gameButtons)

	return retVal
}

func chooseCard(value string, suit string) *canvas.Image {
	image := canvas.NewImageFromFile("gui/svg_playing_cards/fronts/png_96_dpi/" + suit + "_" + value + ".png")
	fmt.Println(image)
	return image
}