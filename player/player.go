package player

import (
	dealer "main/Dealer"
	"main/cards"
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
func PlaceBet(player Player, amount float32) Player {
	if amount > MaxBet {
		amount = MaxBet
	} else if amount < MinBet {
		amount = MinBet
	}
	player.bet = amount
	player.money -= amount
	return player
}

// Close the player's bet and give them twice the money
// They bet if they win, their money back if they tie,
// Or none of it if they lose
func CloseBet(player Player, res result.Result) Player {
	switch res {
	case result.WIN:
		player.money += 2 * player.bet

	case result.TIE:
		player.money += player.bet

	case result.LOSS:
		break
	}

	player.bet = 0
	return player
}

// Get the current bet of the player
func GetBet(player Player) float32 {
	return player.bet
}

// Get the money the player has remaining
func GetMoney(player Player) float32 {
	return player.money
}

// Test if the player has any money left in the game
func HasMoneyLeft(player Player) bool {
	return GetMoney(player) > 0
}

// Get the name of the player
func GetName(player Player) string {
	return player.playerName
}

// Make the player hit the deck
func PlayerHit(player Player, dlr dealer.Dealer) (Player, dealer.Dealer) {
	if !cards.IsBust(player.Hand) && !cards.IsBlackjack(player.Hand) {
		player.Hand = cards.AddCard(dealer.DealCard(dlr), player.Hand)
		player.action = HIT
	} else {
		player.action = STAND
	}
	return player, dlr
}

// Make the given player stand (is done taking hits)
func PlayerStand(player Player) Player {
	player.action = STAND
	return player
}

// Get the current action the player is taking
func GetPlayerAction(player Player) PlayerAction {
	return player.action
}
