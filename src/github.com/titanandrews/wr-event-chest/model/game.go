package model

import (
	"math/rand"
	"time"
)

const (
	// TODO Not sure if this is accurate?
	GoldOpensPerSuperChest = 7

	GoldChestCost = 1000
)

type Game struct {
	Gold                        int
	Prizes                      map[string]int
	GoldOpeningsUntilSuperChest int
	GoldChestsOpened            int
	SuperChestsOpened           int
	RandomGenerator             *rand.Rand
}

func NewGame(startingGold int) *Game {
	s := rand.NewSource(time.Now().UnixNano())
	return &Game{
		Gold:   startingGold,
		Prizes: make(map[string]int),
		GoldOpeningsUntilSuperChest: GoldOpensPerSuperChest,
		RandomGenerator:             rand.New(s),
	}
}
