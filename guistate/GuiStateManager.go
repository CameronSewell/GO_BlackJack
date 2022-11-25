package guistate

import (
	"main/cards"
	"main/result"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
)

var imageWidth float32 = float32(234)
var imageHeight float32 = float32(233)
var cardSize fyne.Size = fyne.Size{Width: imageWidth, Height: imageHeight}

var DealerHand *fyne.Container
var PlayerHand *fyne.Container
var AiPlayersHands []*fyne.Container

var PlayerBet binding.Float = binding.NewFloat()
var PotTotal binding.Float = binding.NewFloat()

var BetString binding.String = binding.FloatToStringWithFormat(PlayerBet, "Bet: %.2f$")
var TotalPotString binding.String = binding.FloatToStringWithFormat(PotTotal, "Money: %.2f$")
var TotalHandString binding.String = binding.NewString()

var PlayerResult result.Result
var PlayerPayout float32
var PlayerHandTotal int

var AIHandTotals []int
var AIPotTotals []float32
var AIPayouts []float32
var AIResults []result.Result

var GameWindow fyne.Window
var GameApp fyne.App

func chooseCard(c cards.Card, isUp bool) *canvas.Image {
	var image = canvas.NewImageFromFile("game/svg_playing_cards/backs/png_96_dpi/blue.png")

	//If it is face up change image from back to the corresponding front of the given card

	if isUp {
		strValue := c.String()
		image = canvas.NewImageFromFile("game/svg_playing_cards/fronts/png_96_dpi/" + strValue + ".png")
	}

	return image
}

func SetCards(h cards.Hand, c *fyne.Container, wait bool) {
	c.RemoveAll()
	var cardImages []*canvas.Image
	var isFaceUp = h.GetFaceUp()
	cardImages = make([]*canvas.Image, h.GetHandCount())

	//For each card in the hands cards get the corresponding card image
	for i, card := range h.GetCards() {

		cardImage := chooseCard(card, isFaceUp[i])
		cardImages[i] = cardImage
	}

	//For each card image set fill mode and add to the container
	for _, e := range cardImages {
		e.FillMode = canvas.ImageFillOriginal //set fill mode for the card image
		e.Resize(cardSize)
		c.Add(e)
		c.Add(layout.NewSpacer())
	}

	//add padding at the end of the new container to center cards
	//c.Add(layout.NewSpacer())

	//set chosen container to the new cards
	if wait {
		time.Sleep(time.Second)
	}
}
