package ai

import (
	"main/dealer"
	"main/guistate"
	"main/player"
	"math/rand"
	"time"
)

// Represents an AI player with an internal
// Player struct and a threshold
type AI struct {
	Plr       player.Player
	threshold float32
}

const (
	MILD       = 0.25
	MODERATE   = 0.5
	AGGRESSIVE = 0.75
)

// Return a new AI struct with its
func NewAI(threshold float32, name string) AI {
	aiPlayer := AI{
		Plr:       player.NewPlayer(name),
		threshold: threshold,
	}
	return aiPlayer
}

// Place a bet for the AI
func (aiPlayer *AI) PlaceBet() {
	bet := player.MinBet + rand.Float32()*(player.MaxBet-player.MinBet)
	aiPlayer.Plr.PlaceBet(bet)
}

// AI player keeps hitting until they choose not to or
// They cannot
func (aiPlayer *AI) AIPlay(dlr *dealer.Dealer, i int) {
	var hit bool
	for hit {
		//Hit if the randomly generated float between 0 and 1
		//Is greater than the threshold
		if rand.Float32() > aiPlayer.threshold {
			aiPlayer.Plr.PlayerHit(dlr, false)
			guistate.SetCards(aiPlayer.Plr.Hand, guistate.AiPlayersHands[i], true)
			time.Sleep(time.Second)
			//Else stand and stop taking hits
		} else {
			hit = false
			aiPlayer.Plr.PlayerStand()
		}
	}
}
