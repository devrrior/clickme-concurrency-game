package main

import (
	"pixel-game-1/scenes"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	mainScene := scenes.NewMainScene()
	pixelgl.Run(mainScene.Run)
}
