package gui

import (
	"fmt"
	"image/color"
	"log"
	"main/game"
	"main/guistate"
	"main/player"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MakeCardContainer() *fyne.Container {
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer())
}

// can set these by saying gui.DealerHand := gui.SetCards(Dealer.Hand)

func GameScreen() *fyne.Container {

	//Set of buttons and container for them
	hitButton := widget.NewButton("Hit", func() {
		if !game.GetPlayer().Hand.IsBust() {
			game.PlayerMove(player.HIT)
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
	standButton := widget.NewButton("Stand", func() { game.PlayerMove(player.STAND) })
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

	guistate.PlayerHand = MakeCardContainer()
	guistate.DealerHand = MakeCardContainer()
	//Player card set up
	cardAlignment := guistate.PlayerHand

	//Dealer card set up
	dealerCardAlignment := guistate.DealerHand

	aiCount := game.GetAICount()
	guistate.AiPlayersHands = make([]*fyne.Container, aiCount)
	for i := 0; i < aiCount; i++ {
		guistate.AiPlayersHands[i] = MakeCardContainer()
	}

	//My attempt at adding AI player's hands with labels to game screen (may need adjustment based on scaling)
	if aiCount == 1 {
		//add player one's label to list of labels going across horizontally
		handLabelAlignment = container.NewHBox(playerHandLabel, layout.NewSpacer(), aiOneLabelAlignment)
		//add player one's cards to horizontal alignment of card containers
		cardAlignment = container.New(layout.NewHBoxLayout(), guistate.PlayerHand, layout.NewSpacer(), guistate.AiPlayersHands[0])
	} else if aiCount == 2 {
		handLabelAlignment = container.NewHBox(playerHandLabel, layout.NewSpacer(), aiOneLabelAlignment, layout.NewSpacer(),
			aiTwoLabelAlignment)
		cardAlignment = container.New(layout.NewHBoxLayout(), guistate.PlayerHand, layout.NewSpacer(), guistate.AiPlayersHands[0],
			layout.NewSpacer(), guistate.AiPlayersHands[1])
	}

	//combines labels with card containers
	labelAlignment := container.New(layout.NewVBoxLayout(), betLabelAlignment, totalPotLabelAlignment,
		dealerLabelAlignment, dealerCardAlignment, handLabelAlignment, cardAlignment)

	//Putting it all together
	retVal := container.New(layout.NewVBoxLayout(), labelAlignment, layout.NewSpacer(), layout.NewSpacer(), gameButtons)

	return retVal
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
