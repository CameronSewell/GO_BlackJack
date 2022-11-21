package gui

import (
	"fmt"
	"image/color"
	"log"
	"main/Game"
	"main/cards"
	"main/player"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//can set these by saying gui.DealerHand := gui.SetCards(Dealer.Hand)
var DealerHand *fyne.Container
var PlayerHand *fyne.Container
var AiPlayersHands []*fyne.Container

func GameScreen() *fyne.Container {

	//Set of buttons and container for them
	hitButton := widget.NewButton("Hit", func() {
		if !Game.GetPlayer().Hand.IsBust() {
			Game.PlayerMove(player.HIT)
		} else {
			fmt.Println("Your Hand is busted. Cannot hit again.")
		}
	})
	splitButton := widget.NewButton("Split", func() { log.Println("split button tapped") })
	doubleButton := widget.NewButton("Double", func() { log.Println("double button tapped") })
	surrenderButton := widget.NewButton("Surrender", func() {
		log.Println("surrender button tapped")
		EndScreen()
	})
	standButton := widget.NewButton("Stand", func() { Game.PlayerMove(player.STAND) })
	gameButtons := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), layout.NewSpacer(), hitButton, surrenderButton, doubleButton, splitButton,
		standButton, layout.NewSpacer(), layout.NewSpacer())

	//Game's labels
	betLabel := widget.NewLabel("Bet: ")
	totalPotLabel := widget.NewLabel("$: ")
	playerHandLabel := widget.NewLabel("Your hand: ")
	dealerHandLabel := widget.NewLabel("Dealer's hand: ")
	aiOneHandLabel := widget.NewLabel("Player 1's hand: ")
	aiTwoHandLabel := widget.NewLabel("Player 2's hand: ")

	//Aligning game labels left
	betLabelAlignment := container.NewHBox(betLabel, layout.NewSpacer(), layout.NewSpacer())
	totalPotLabelAlignment := container.NewHBox(totalPotLabel, layout.NewSpacer(), layout.NewSpacer())
	handLabelAlignment := container.NewHBox(layout.NewSpacer(), playerHandLabel, layout.NewSpacer())
	dealerLabelAlignment := container.NewHBox(layout.NewSpacer(), dealerHandLabel, layout.NewSpacer())
	aiOneLabelAlignment := container.NewHBox(layout.NewSpacer(), aiOneHandLabel, layout.NewSpacer())
	aiTwoLabelAlignment := container.NewHBox(layout.NewSpacer(), aiTwoHandLabel, layout.NewSpacer())

	//Player card set up
	cardAlignment := PlayerHand

	//Dealer card set up
	dealerCardAlignment := DealerHand

	//My attempt at adding AI player's hands with labels to game screen (may need adjustment based on scaling)
	if len(AiPlayersHands) == 1 {
		//add player one's label to list of labels going across horizontally
		handLabelAlignment = container.NewHBox(playerHandLabel, layout.NewSpacer(), aiOneLabelAlignment)
		//add player one's cards to horizontal alignment of card containers
		cardAlignment = container.New(layout.NewHBoxLayout(), PlayerHand, layout.NewSpacer(), AiPlayersHands[0])
	} else if len(AiPlayersHands) == 2 {
		handLabelAlignment = container.NewHBox(playerHandLabel, layout.NewSpacer(), aiOneLabelAlignment, layout.NewSpacer(),
			aiTwoLabelAlignment)
		cardAlignment = container.New(layout.NewHBoxLayout(), PlayerHand, layout.NewSpacer(), AiPlayersHands[0],
			layout.NewSpacer(), AiPlayersHands[1])
	}

	//combines labels with card containers
	labelAlignment := container.New(layout.NewVBoxLayout(), betLabelAlignment, totalPotLabelAlignment,
		dealerLabelAlignment, dealerCardAlignment, handLabelAlignment, cardAlignment)

	//Putting it all together
	retVal := container.New(layout.NewVBoxLayout(), labelAlignment, layout.NewSpacer(), layout.NewSpacer(), gameButtons)

	return retVal
}

func chooseCard(value string, suit string, isUp bool) *canvas.Image {
	var image = canvas.NewImageFromFile("gui/svg_playing_cards/backs/png_96_dpi/blue.png")

	//If it is face up change image from back to the corresponding front of the given card
	if isUp {
		image = canvas.NewImageFromFile("gui/svg_playing_cards/fronts/png_96_dpi/" + suit + "_" + value + ".png")
	}

	return image
}

func SetCards(h cards.Hand, c *fyne.Container) {
	var cardImages []*canvas.Image
	var isFaceUp = h.GetFaceUp()

	//For each card in the hands cards get the corresponding card image
	for i, card := range h.GetCards() {
		suit := card.Suit
		value := card.Value
		cardImage := chooseCard(strconv.Itoa(value), suit, isFaceUp[i])
		cardImages[i] = cardImage
	}

	//add padding at the beginning of the new container to center cards
	cardAlignment := container.New(layout.NewHBoxLayout(), layout.NewSpacer())

	//For each card image set fill mode and add to the container
	for _, e := range cardImages {
		e.FillMode = canvas.ImageFillOriginal //set fill mode for the card image
		cardAlignment.Add(e)
	}

	//add padding at the end of the new container to center cards
	cardAlignment.Add(layout.NewSpacer())

	//set chosen container to the new cards
	c = cardAlignment
}

func EndScreen() *fyne.Container {
	gameOverText := canvas.NewText("Game Over", color.White)
	gameOverText.TextSize = 50
	gameOverAlign := container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), gameOverText, layout.NewSpacer(), layout.NewSpacer())
	newGameButton := widget.NewButton("New Game", func() {
		log.Println("double button tapped")
		StartScreen()
	})
	endButton := widget.NewButton("Exit", func() {
		log.Println("split button tapped")
		os.Exit(0)
	})

	choiceButtons := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), newGameButton, endButton, layout.NewSpacer())

	retVal := container.New(layout.NewVBoxLayout(), gameOverAlign, layout.NewSpacer(), choiceButtons, layout.NewSpacer(), layout.NewSpacer())

	return retVal
}
