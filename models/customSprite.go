package models

import (
	"image"
	"math/rand"
	"os"

	"github.com/faiface/pixel"
)

type CustomSprite struct {
	Sprite  *pixel.Sprite
	Matrix  pixel.Matrix
	Clicked bool
}

func NewCustomSprite(sprite *pixel.Sprite, matrix pixel.Matrix) *CustomSprite {
	return &CustomSprite{
		Sprite:  sprite,
		Matrix:  matrix,
		Clicked: false,
	}
}

func CreateRandomSprite(bounds pixel.Rect) *CustomSprite {
	pic, err := loadPicture("assets/images/ghost.png")

	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	matrix := pixel.IM
	matrix = matrix.ScaledXY(pixel.ZV, pixel.V(.2, .2))
	matrix = matrix.Moved(pixel.V(rand.Float64()*bounds.W(), rand.Float64()*bounds.H()))

	return NewCustomSprite(sprite, matrix)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
