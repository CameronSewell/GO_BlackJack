package guistate

import (
	"main/cards"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
)

var imageWidth float32 = float32(234)
var imageHeight float32 = float32(233)
var cardSize fyne.Size = fyne.Size{Width: imageWidth, Height: imageHeight}

var DealerHand *fyne.Container
var PlayerHand *fyne.Container
var AiPlayersHands []*fyne.Container

var BetString binding.String
var TotalPotString binding.String
var TotalHandString binding.String

var GameWindow fyne.Window

func chooseCard(c cards.Card, isUp bool) *canvas.Image {
	var image = canvas.NewImageFromFile("gui/svg_playing_cards/backs/png_96_dpi/blue.png")

	//If it is face up change image from back to the corresponding front of the given card

	if isUp {
		strValue := c.String()
		image = canvas.NewImageFromFile("gui/svg_playing_cards/fronts/png_96_dpi/" + strValue + ".png")
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
	}

	//add padding at the end of the new container to center cards
	//c.Add(layout.NewSpacer())

	//set chosen container to the new cards
	if wait {
		time.Sleep(750 * time.Millisecond)
	}
}
