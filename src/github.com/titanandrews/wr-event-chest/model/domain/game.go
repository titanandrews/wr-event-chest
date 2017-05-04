package domain

import (
	"fmt"
	"strconv"

	"math/rand"

	"github.com/titanandrews/wr-event-chest/model"
)

type Game interface {
	OpenChest()
	Show()
}

type game struct {
	game *model.Game
}

func NewGame(startingGold int) Game {
	return &game{model.NewGame(startingGold)}
}

func (g *game) OpenChest() {
	var chestPrizes []string
	if g.game.GoldOpeningsUntilSuperChest <= 0 {
		fmt.Println("Opening super chest...")
		// Reset
		g.game.GoldOpeningsUntilSuperChest = model.GoldOpensPerSuperChest
		chestPrizes = g.populateChest(model.SuperChest)
		g.game.SuperChestsOpened += 1

	} else {
		if g.game.Gold < model.GoldChestCost {
			fmt.Println("Not enough gold to open the chest.")
			return
		}

		fmt.Println("Opening gold chest...")
		g.game.Gold -= model.GoldChestCost
		g.game.GoldOpeningsUntilSuperChest -= 1
		g.game.GoldChestsOpened += 1
		chestPrizes = g.populateChest(model.GoldChest)
	}

	chestPrizes = g.shufflePrizes(chestPrizes)
	i := g.game.RandomGenerator.Intn(len(chestPrizes) - 1)
	wonPrize := chestPrizes[i]
	gold, err := strconv.Atoi(wonPrize)
	if err == nil {
		wonPrize = fmt.Sprintf("%s Au", wonPrize)
		g.game.Gold += gold
	}

	fmt.Println("You won", wonPrize)
	fmt.Println("Current gold:", g.game.Gold)
	g.game.Prizes[wonPrize] += 1
}

func (g game) shufflePrizes(prizes []string) []string {
	dest := make([]string, len(prizes))
	perm := rand.Perm(len(prizes))
	for i, v := range perm {
		dest[v] = prizes[i]
	}

	return dest
}

func (g game) populateChest(chest map[string]int) []string {
	// Fill the chest with each prize the same number as the drop rate.
	var prizes []string
	for p, dr := range chest {
		for i := 0; i < dr; i++ {
			prizes = append(prizes, p)
		}
	}

	return prizes
}

func (g *game) Show() {
	fmt.Println("--- Current winnings ---")
	fmt.Println("Current gold:", g.game.Gold)
	fmt.Println("Gold chests opened:", g.game.GoldChestsOpened)
	fmt.Println("Super chests opened:", g.game.SuperChestsOpened)
	fmt.Println("Prizes:")
	for p, c := range g.game.Prizes {
		fmt.Println(p, c)
	}
}
