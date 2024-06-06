package main

import (
	//"github.com/hajimehoshi/ebiten" usar a v2
	"jogo-2d-golang/game"

	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	g := game.NewGame()

	err := ebiten.RunGame(g)
	if (err != nil) {
		panic(err)
	}


}