package Game

import (
	"main/cards"
	"main/player"
	"main/result"
	"math/rand"
)

type GameState int

const (
	GAME_START GameState = iota
	PLAYER_TURN
	DEALER_TURN
	GAME_END
)

// Struct containing data for managing the game
type GameManager struct {
	dealer  Dealer
	players []player.Player
	state   GameState
	//ai     []AI
}

// Contains the current state of the running game
var gm GameManager

// Create a new game and initialize all players
// TODO: add AI data
func NewGame(names []string) {
	gm = GameManager{}
	for i := 0; i < len(names); i++ {
		gm.players[i] = player.NewPlayer(names[i])
	}
}

// Create a new game and initialize everyone
// TODO: add AI functionality
func StartGame(bet float32) {
	//Initialize new dealer
	gm.dealer = NewDealer()

	//Place new bets for everyone and remake starting hands
	for i := 0; i < len(gm.players); i++ {
		//Randomly generate bet if player is AI (position greater than zero)
		if i > 0 {
			bet = player.MinBet + rand.Float32()*(player.MaxBet-player.MinBet)
		}

		//Place bet
		gm.players[i] = player.PlaceBet(gm.players[i], bet)

		//Create a starting hand
		gm.players[i].Hand = DealStartingHand(gm.dealer, gm.players[i].Hand)
	}

	//Reset the game state to start
	gm.state = GAME_START

	//Deal starting hands
	gm.dealer.Hand = DealStartingHand(gm.dealer, gm.dealer.Hand)
}

// Play the game
func PlayGame() {

	gm.state = PLAYER_TURN

	//Continuously run the game until it ends
	for gm.state != GAME_END {
		switch gm.state {
		//Player turn in game (goes first)
		case PLAYER_TURN:
			PlayGamePlayers(player.START)

		//Let the dealer play when it's their turn
		case DEALER_TURN:
			DealerPlay(gm.dealer)
			gm.state = GAME_END

		}
	}

	//Finish running the current blackjack game
	EndGame()
}

// Prompt and respond to actions from the user
func PlayGamePlayers(action player.PlayerAction) {
	for i := 0; i < len(gm.players); i++ {
		for action != player.STAND {
			//TODO: figure out AI logic VS player logic
			//For hitting and standing
			action = player.STAND
		}
	}
}

// End the running Blackjack game for all players
func EndGame() {
	//Loop over all players in the game
	for i := 0; i < len(gm.players); i++ {
		//Get the current player
		p := GetPlayer(i)
		var r result.Result

		//Check if the player won, lost, or tied with the dealer

		//If the player busted, it means they lost
		if cards.IsBust(p.Hand) {
			r = result.LOSS

			//Else, check against all other conditions
		} else {

			//If the dealer busted and the player didn't, the player wins the hand
			if cards.IsBust(gm.dealer.Hand) {
				r = result.WIN

				//Else compare the dealer's final total against the player's
			} else {
				//Get the totals
				dealerTotal := cards.GetHandTotal(gm.dealer.Hand)
				playerTotal := cards.GetHandTotal(p.Hand)

				//The player won if their result was greater than the dealer's total
				if playerTotal > dealerTotal {
					r = result.WIN

					//The player tied if their total is equal to the dealer's
				} else if playerTotal == dealerTotal {
					r = result.TIE

					//Else, they lost the match
				} else {
					r = result.LOSS
				}
			}
		}

		//Update the bets of the
		gm.players[i] = player.CloseBet(gm.players[i], r)
	}
}

// Get the number of players playing the game
func GetPlayerCount() int {
	return len(gm.players)
}

// Get the player at the specified index
func GetPlayer(i int) player.Player {
	return gm.players[i]
}
