package ai

import (
	dealer "main/Dealer"
	"main/player"
	"math/rand"
)

// Represents an AI player with an internal
// Player struct and a threshold
type AI struct {
	Plr       player.Player
	threshold float32
}

// Return a new AI struct with its
func NewAI(threshold float32, name string) AI {
	aiPlayer := AI{
		Plr:       player.NewPlayer(name),
		threshold: threshold,
	}
	return aiPlayer
}

// Place a bet for the AI
func PlaceBet(aiPlayer AI) AI {
	bet := player.MinBet + rand.Float32()*(player.MaxBet-player.MinBet)
	aiPlayer.Plr = player.PlaceBet(aiPlayer.Plr, bet)
	return aiPlayer
}

// AI player keeps hitting until they choose not to or
// They cannot
func AIPlay(aiPlayer AI, dlr dealer.Dealer) AI {
	var hit bool
	for hit {
		//Hit if the randomly generated float between 0 and 1
		//Is greater than the threshold
		if rand.Float32() > aiPlayer.threshold {
			player.PlayerHit(aiPlayer.Plr, dlr)
			//Else stand and stop taking hits
		} else {
			hit = false
			player.PlayerStand(aiPlayer.Plr)
		}
	}
	return aiPlayer
}
