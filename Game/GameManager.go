package game

import ("fmt"
	dealer "main/Dealer"
	"main/ai"
	"main/cards"
	"main/guistate"
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

func GetState() GameState {
	return gm.state
}

// Create a new game and initialize all players
func NewGame(names []string, thresholds []float32) {
	gm = GameManager{}
	gm.player = player.NewPlayer(names[0])
	gm.aiPlayers = make([]ai.AI, len(thresholds))
	for i := 0; i < len(thresholds); i++ {
		fmt.Println(thresholds[i])
		fmt.Println(names[i+1])
		gm.aiPlayers[i] = ai.NewAI(thresholds[i], names[i+1])
	}
	gm.state = GAME_START
}

// Create a new game and initialize everyone
func StartGame(bet float32) {
	//Initialize new dealer
	if gm.state == GAME_START {
		gm.dlr = dealer.NewDealer()
		gm.player.PlaceBet(bet)
		gm.dlr.DealStartingHand(&gm.player.Hand, true)
		guistate.SetCards(gm.player.Hand, guistate.PlayerHand, true)

		//Place new bets for everyone and remake starting hands
		for i := 0; i < len(gm.aiPlayers); i++ {
			//Randomly generate bet if player is AI (position greater than zero)

			//Place bet
			gm.aiPlayers[i].PlaceBet()

			//Create a starting hand
			gm.dlr.DealStartingHand(&gm.aiPlayers[i].Plr.Hand, false)
			guistate.SetCards(gm.aiPlayers[i].Plr.Hand, guistate.AiPlayersHands[i], true)
		}

		//Deal starting hands
		gm.dlr.DealStartingHand(&gm.dlr.Hand, false)
		gm.dlr.Hand.SetFirstUp()
		guistate.SetCards(gm.dlr.Hand, guistate.DealerHand, true)

		//Reset the game state to start
		gm.state = AI_TURN
		AIMoves()
	}
}

// TODO: Figure out how to connect this to the frontend
func PlayerMove(action player.PlayerAction) {
	if gm.state == PLAYER_TURN {
		switch action {
		case player.HIT:
			gm.player.PlayerHit(&gm.dlr, true)
			guistate.SetCards(gm.player.Hand, guistate.PlayerHand, true)
		case player.STAND:
			gm.player.PlayerStand()
			gm.state = DEALER_TURN
			DealerMoves()
		}
	}
}

// Make each of the AIs play the game until they're done
func AIMoves() {
	if gm.state == AI_TURN {
		for i := 0; i < len(gm.aiPlayers); i++ {
			gm.aiPlayers[i].AIPlay(&(gm.dlr), i)
		}
		gm.state = PLAYER_TURN
	}
}

// The dealer makes their move
func DealerMoves() {
	if gm.state == DEALER_TURN {
		gm.dlr.DealerPlay()
		gm.state = GAME_END
		EndGame()
	}
}

// Get the result of the given hand by comparing it
// Against the dealer's
func getResult(h cards.Hand) result.Result {
	//If the player busted, it means they lost
	var r result.Result
	if h.IsBust() {
		r = result.LOSS

		//Else, check against all other conditions
	} else {

		//If the dealer busted and the player didn't, the player wins the hand
		if gm.dlr.Hand.IsBust() {
			r = result.WIN

			//Else compare the dealer's final total against the player's
		} else {
			//Get the totals
			dealerTotal := gm.dlr.Hand.GetHandTotal()
			playerTotal := h.GetHandTotal()

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
	if gm.state == GAME_END {
		var r result.Result
		gm.dlr.Hand.SetUp()
		guistate.SetCards(gm.dlr.Hand, guistate.DealerHand, false)
		for i := 0; i < len(gm.aiPlayers); i++ {
			gm.aiPlayers[i].Plr.Hand.SetUp()
			guistate.SetCards(gm.aiPlayers[i].Plr.Hand, guistate.AiPlayersHands[i], false)
			r = getResult(gm.aiPlayers[i].Plr.Hand)

			//Update the bets of the AI players
			gm.aiPlayers[i].Plr.CloseBet(r)
		}
		r = getResult(gm.player.Hand)
		gm.player.Hand.SetUp()
		guistate.SetCards(gm.player.Hand, guistate.PlayerHand, false)
		gm.player.CloseBet(r)
	}
}

// Get the number of players playing the game
func GetAICount() int {
	return len(gm.aiPlayers)
}

// Get the player at the specified index
func GetPlayer() *player.Player {
	return &gm.player
}

// Get the AI at the specific index
func GetAI(i int) ai.AI {
	return gm.aiPlayers[i]
}
