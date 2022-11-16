package player

import (
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

const ()

// Defines a player struct
type Player struct {
	Hand       cards.Hand
	playerName string
	money      float32
	bet        float32
}

// Create a new player with a hand and starting bet
// of zero
func NewPlayer(name string) Player {
	p := Player{
		Hand:       cards.NewHand(),
		playerName: name,
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
// They bet if they win
func CloseBet(player Player, res result.Result) Player {
	switch res {
	case result.WIN:
		player.money += 2 * player.bet
		break
	case result.TIE:
		player.money += player.bet
		break
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

func GetName(player Player) string {
	return player.playerName
}
