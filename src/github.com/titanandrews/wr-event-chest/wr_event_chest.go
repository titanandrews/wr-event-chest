package main

import (
	"strconv"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/titanandrews/wr-event-chest/model/domain"
)

var game domain.Game

func main() {
	shell := ishell.New()

	shell.Println("War Robots Event Chest Simulator")
	shell.Println("Type help for instructions")

	shell.AddCmd(&ishell.Cmd{
		Name: "new",
		Help: "new game",
		Func: func(c *ishell.Context) {
			c.Println("How much gold would you like to start with?")
			input := c.ReadLine()
			input = strings.TrimSpace(input)
			gold, err := strconv.Atoi(input)
			if err != nil {
				c.Println("Not a number")
				return
			}

			game = domain.NewGame(gold)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "open",
		Help: "open chest",
		Func: func(c *ishell.Context) {
			if game == nil {
				c.Println("A game has not yet been started. Type new to start a new game.")
				return
			}

			game.OpenChest()
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "show",
		Help: "show current winnings",
		Func: func(c *ishell.Context) {
			if game == nil {
				c.Println("A game has not yet been started. Type new to start a new game.")
				return
			}

			game.Show()
		},
	})

	shell.Start()
}
