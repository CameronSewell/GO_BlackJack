package game

import (
	"log"
	"main/guistate"
	"main/player"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func showBetPopup() {
	var modal *widget.PopUp
	input := widget.NewEntry()
	modal = widget.NewModalPopUp(container.New(
		layout.NewHBoxLayout(),
		widget.NewLabel("Enter your bet:"),
		input,
		widget.NewButton("Enter", func() {
			modal.Hide()
			if bet, err := strconv.ParseFloat(input.Text, 32); err == nil {
				StartGame(float32(bet))
			} else {
				panic(err)
			}
		}),
	),
		guistate.GameWindow.Canvas())
	modal.Show()
}

func showErrorPopup(message string) {
	var errorModal *widget.PopUp
	errorModal = widget.NewModalPopUp(container.New(
		layout.NewVBoxLayout(),
		widget.NewLabel(message),
		widget.NewButton("Close", func() {
			errorModal.Hide()
		}),
	),
		guistate.GameWindow.Canvas())
	errorModal.Show()
}

func MakeCardContainer() *fyne.Container {
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer())
}

func GameScreen() {
	//Set of buttons and container for them
	hitButton := widget.NewButton("Hit", func() {
		if !GetPlayer().Hand.IsBust() {
			PlayerMove(player.HIT)
		} else {
			showErrorPopup("You hand is busted. You must stand.")
		}
	})
	splitButton := widget.NewButton("Split", func() { log.Println("split button tapped") })
	doubleButton := widget.NewButton("Double", func() {
		if !GetPlayer().Hand.IsBust() {
			PlayerMove(player.HIT)
		} else {
			showErrorPopup("You hand is busted. You must stand.")
		}
	})
	surrenderButton := widget.NewButton("Surrender", func() {
		log.Println("surrender button tapped")
		StartScreen()
	})
	standButton := widget.NewButton("Stand", func() { PlayerMove(player.STAND) })
	gameButtons := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), layout.NewSpacer(), hitButton, surrenderButton, doubleButton, splitButton,
		standButton, layout.NewSpacer(), layout.NewSpacer())

	//Game's labels
	betLabel := widget.NewLabelWithData(guistate.BetString)
	totalPotLabel := widget.NewLabelWithData(guistate.TotalPotString)
	playerHandLabel := widget.NewLabelWithData(guistate.TotalHandString)
	dealerHandLabel := widget.NewLabel("Dealer's Hand:")
	aiOneHandLabel := widget.NewLabel("Player 1's Hand: ")
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

	aiCount := GetAICount()
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

	guistate.GameWindow.SetContent(retVal)
	showBetPopup()
}
