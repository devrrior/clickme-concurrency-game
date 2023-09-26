package scenes

import (
	"fmt"
	"image"
	"math/rand"
	"os"
	"pixel-game-1/models"
	"sync"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type MainScene struct {
}

func NewMainScene() *MainScene {
	return &MainScene{}
}

var (
	sprites []models.CustomSprite
	mutex   sync.Mutex
	counter int
	timer   int
)

func handleLogic(bounds pixel.Rect) {
	for {
		time.Sleep(1 * time.Second)

		mutex.Lock()
		sprites = []models.CustomSprite{}
		sprites = append(sprites, createRandomSprite(bounds))
		mutex.Unlock()
	}
}

func handleInput(win *pixelgl.Window) {
	for !win.Closed() {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			pos := win.MousePosition()

			mutex.Lock()
			for i := range sprites {
				rect := sprites[i].Sprite.Frame().Moved(sprites[i].Matrix.Project(pixel.ZV))
				if rect.Contains(pos) && !sprites[i].Clicked {
					sprites[i].Clicked = true
					counter++
				}
			}
			mutex.Unlock()
		}

		time.Sleep(1 * time.Millisecond)
	}
}

func handleTimer() {
	for {
		time.Sleep(1 * time.Second)

		mutex.Lock()
		timer++
		if timer >= 60 {
			timer = 0
			counter = 0
		}
		mutex.Unlock()
	}
}

func createRandomSprite(bounds pixel.Rect) models.CustomSprite {
	pic, err := loadPicture("assets/images/ghost.png")

	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	matrix := pixel.IM
	matrix = matrix.ScaledXY(pixel.ZV, pixel.V(.2, .2))
	matrix = matrix.Moved(pixel.V(rand.Float64()*bounds.W(), rand.Float64()*bounds.H()))

	return models.NewCustomSprite(sprite, matrix)
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

func (s *MainScene) Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	go handleLogic(win.Bounds())
	go handleInput(win)
	go handleTimer()

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(50, 50), atlas)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		mutex.Lock()
		for i := range sprites {
			if !sprites[i].Clicked {
				sprites[i].Sprite.Draw(win, sprites[i].Matrix)
			}
		}

		txt.Clear()
		fmt.Fprintf(txt, "Counter: %d\nTimer: %02d:%02d", counter, timer/60, timer%60)
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))

		mutex.Unlock()

		win.Update()
	}
}
