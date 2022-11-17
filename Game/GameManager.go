package Game

import (
	dealer "main/Dealer"
	"main/ai"
	"main/cards"
	"main/player"
	"main/result"
)

// Represents the game state
type GameState int

// Enum of possible game states
const (
	GAME_START GameState = iota
	PLAYER_TURN
	AI_TURN
	DEALER_TURN
	GAME_END
)

// Struct containing data for managing the game
type GameManager struct {
	dlr       dealer.Dealer
	player    player.Player
	state     GameState
	aiPlayers []ai.AI
}

// Contains the current state of the running game
var gm GameManager

// Create a new game and initialize all players
func NewGame(names []string, thresholds []float32) {
	gm = GameManager{}
	gm.player = player.NewPlayer(names[0])
	for i := 0; i < len(thresholds); i++ {
		gm.aiPlayers[i] = ai.NewAI(thresholds[i], names[i+1])
	}
}

// Create a new game and initialize everyone
func StartGame(bet float32) {
	//Initialize new dealer
	gm.dlr = dealer.NewDealer()
	gm.player = player.PlaceBet(gm.player, bet)
	gm.player.Hand = dealer.DealStartingHand(gm.dlr, gm.player.Hand)

	//Place new bets for everyone and remake starting hands
	for i := 0; i < len(gm.aiPlayers); i++ {
		//Randomly generate bet if player is AI (position greater than zero)

		//Place bet
		gm.aiPlayers[i] = ai.PlaceBet(gm.aiPlayers[i])

		//Create a starting hand
		gm.aiPlayers[i].Plr.Hand = dealer.DealStartingHand(gm.dlr, gm.aiPlayers[i].Plr.Hand)
	}

	//Deal starting hands
	gm.dlr.Hand = dealer.DealStartingHand(gm.dlr, gm.dlr.Hand)

	//Reset the game state to start
	gm.state = GAME_START
}

// Play the game
func PlayGame() {

	gm.state = PLAYER_TURN

	//Continuously run the game until it ends
	for gm.state != GAME_END {
		switch gm.state {
		//AI turns in game (AIs go first)
		case AI_TURN:
			playGameAIs()
			gm.state = PLAYER_TURN
		//Player goes second to last
		case PLAYER_TURN:
			PlayGamePlayers(player.START)
			gm.state = DEALER_TURN
		//Let the dealer play when it's their turn
		case DEALER_TURN:
			gm.dlr = dealer.DealerPlay(gm.dlr)
			gm.state = GAME_END

		}
	}

	//Finish running the current blackjack game
	EndGame()
}

// TODO: Figure out how to connect this to the frontend
func PlayGamePlayers(action player.PlayerAction) {
	//TODO: put this in the frontend somehow
}

// Make each of the AIs play the game until they're done
func playGameAIs() {
	for i := 0; i < len(gm.aiPlayers); i++ {
		gm.aiPlayers[i] = ai.AIPlay(gm.aiPlayers[i], gm.dlr)
	}
}

// Get the result of the given hand by comparing it
// Against the dealer's
func getResult(h cards.Hand) result.Result {
	//If the player busted, it means they lost
	var r result.Result
	if cards.IsBust(h) {
		r = result.LOSS

		//Else, check against all other conditions
	} else {

		//If the dealer busted and the player didn't, the player wins the hand
		if cards.IsBust(gm.dlr.Hand) {
			r = result.WIN

			//Else compare the dealer's final total against the player's
		} else {
			//Get the totals
			dealerTotal := cards.GetHandTotal(gm.dlr.Hand)
			playerTotal := cards.GetHandTotal(h)

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
	return r
}

// End the running Blackjack game for all players
func EndGame() {
	//Loop over all players in the game
	var r result.Result
	for i := 0; i < len(gm.aiPlayers); i++ {
		r = getResult(gm.aiPlayers[i].Plr.Hand)

		//Update the bets of the
		gm.aiPlayers[i].Plr = player.CloseBet(gm.aiPlayers[i].Plr, r)
	}
	r = getResult(gm.player.Hand)
	gm.player = player.CloseBet(gm.player, r)
}

// Get the number of players playing the game
func GetAICount() int {
	return len(gm.aiPlayers) + 1
}

// Get the player at the specified index
func GetPlayer() player.Player {
	return gm.player
}

// Get the AI at the specific index
func GetAI(i int) ai.AI {
	return gm.aiPlayers[i]
}
