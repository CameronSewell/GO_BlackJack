package game

import (
	"fmt"
	"image/color"
	"main/guistate"
	"main/result"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func EndScreen() {

	potTotal, err := guistate.PotTotal.Get()
	if err != nil {
		panic(err)
	}

	gameOverText := canvas.NewText("Game Over", color.White)
	gameOverText.TextSize = 50
	gameOverAlign := container.NewHBox(layout.NewSpacer(), layout.NewSpacer(), gameOverText, layout.NewSpacer(), layout.NewSpacer())
	newGameButton := widget.NewButton("Keep Playing", func() {
		if potTotal > 0 {
			GameScreen()
		} else {
			ShowErrorPopup("You are out of money and cannot keep playing.")
		}
	})
	endButton := widget.NewButton("Quit To Start Screen", func() {
		StartScreen()
	})

	var outcome *widget.Label
	var potResult *widget.Label
	var handResult *widget.Label
	var payoutResult *widget.Label
	if guistate.PlayerResult == result.WIN {
		outcome = widget.NewLabel("You won!")
	} else if guistate.PlayerResult == result.LOSS {
		outcome = widget.NewLabel("You lost.")
	} else if guistate.PlayerResult == result.SURRENDER {
		outcome = widget.NewLabel("You surrendered your hand.")
	} else {
		outcome = widget.NewLabel("You tied!")
	}

	handResult = widget.NewLabel(fmt.Sprintf("Your hand value was %d", GetPlayer().Hand.GetHandTotal()))
	payoutResult = widget.NewLabel(fmt.Sprintf("Your payout is %.2f$", guistate.PlayerPayout))
	potResult = widget.NewLabel(fmt.Sprintf("Your money is now %.2f$", GetPlayer().GetMoney()))

	playerResultsContainer := container.New(layout.NewVBoxLayout(), outcome, handResult, payoutResult, potResult)

	aiResultsContainer := container.New(layout.NewVBoxLayout())
	for i := 0; i < GetAICount(); i++ {

		if guistate.AIResults[i] == result.WIN {
			outcome = widget.NewLabel(fmt.Sprintf("Player %d won!", i+1))
		} else if guistate.AIResults[i] == result.LOSS {
			outcome = widget.NewLabel(fmt.Sprintf("Player %d lost.", i+1))
		} else {
			outcome = widget.NewLabel(fmt.Sprintf("Player %d tied!", i+1))
		}

		handResult = widget.NewLabel(fmt.Sprintf("Player's %d hand value is %d", i+1, GetAIPlayer(i).Plr.Hand.GetHandTotal()))
		payoutResult = widget.NewLabel(fmt.Sprintf("Player %d's payout is %.2f$", i+1, guistate.AIPayouts[i]))
		potResult = widget.NewLabel(fmt.Sprintf("Player %d's money is now %.2f$", i+1, GetAIPlayer(i).Plr.GetMoney()))
		aiResultsContainer.Add(container.New(layout.NewVBoxLayout(), outcome, handResult, payoutResult, potResult))
	}

	choiceButtons := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), newGameButton, endButton, layout.NewSpacer())

	retVal := container.New(layout.NewVBoxLayout(), gameOverAlign, layout.NewSpacer(),
		playerResultsContainer, layout.NewSpacer(), aiResultsContainer, layout.NewSpacer(), choiceButtons, layout.NewSpacer())
	guistate.GameWindow.SetContent(retVal)
}
