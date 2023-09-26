package models

import "github.com/faiface/pixel"

type CustomSprite struct {
	Sprite  *pixel.Sprite
	Matrix  pixel.Matrix
	Clicked bool
}

func NewCustomSprite(sprite *pixel.Sprite, matrix pixel.Matrix) CustomSprite {
	return CustomSprite{
		Sprite:  sprite,
		Matrix:  matrix,
		Clicked: false,
	}
}
