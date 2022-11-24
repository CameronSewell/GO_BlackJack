package player

import (
	"main/cards"
	"main/dealer"
	"main/result"
)

// Amount that all players start with
var StartingAmount float32 = 200

// Max bet that a player can make
var MaxBet float32 = 50
var MinBet float32 = 5

// Defines actions that a player can take during the game
type PlayerAction int

const (
	START PlayerAction = iota
	HIT
	STAND
)

// Defines a player struct
type Player struct {
	Hand       cards.Hand
	playerName string
	action     PlayerAction
	money      float32
	bet        float32
}

// Create a new player with a hand and starting bet
// of zero
func NewPlayer(name string) Player {
	p := Player{
		Hand:       cards.NewHand(),
		playerName: name,
		action:     START,
		money:      StartingAmount,
		bet:        0,
	}
	return p
}

// Place the given bet
func (p *Player) PlaceBet(amount float32) {
	if amount > MaxBet {
		amount = MaxBet
	} else if amount < MinBet {
		amount = MinBet
	}
	p.bet = amount
	p.money -= amount
}

// Close the player's bet and give them twice the money
// They bet if they win, their money back if they tie,
// Or none of it if they lose
func (p *Player) CloseBet(res result.Result) {
	switch res {
	case result.WIN:
		p.money += 2 * p.bet

	case result.TIE:
		p.money += p.bet

	case result.LOSS:
		break
	}

	p.bet = 0
}

// Get the current bet of the player
func (p *Player) GetBet() float32 {
	return p.bet
}

// Get the money the player has remaining
func (p *Player) GetMoney() float32 {
	return p.money
}

// Test if the player has any money left in the game
func (p *Player) HasMoneyLeft() bool {
	return p.GetMoney() > 0
}

// Get the name of the player
func (p *Player) GetName() string {
	return p.playerName
}

// Make the player hit the deck
func (p *Player) PlayerHit(dlr *dealer.Dealer, isUp bool) {
	if !p.Hand.IsBust() && !p.Hand.IsBlackjack() {
		p.Hand.AddCard(dlr.DealCard(), isUp)
		p.action = HIT
	} else {
		p.action = STAND
	}
}

// Make the given player stand (is done taking hits)
func (p *Player) PlayerStand() {
	p.action = STAND
}

// Get the current action the player is taking
func (p *Player) GetPlayerAction() PlayerAction {
	return p.action
}
