package scenes

import (
	"fmt"
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
	sprite  models.CustomSprite
	mutex   sync.Mutex
	counter int
	timer   int
)

func handleLogic(bounds pixel.Rect) {
	for {
		time.Sleep(1 * time.Second)

		mutex.Lock()
		sprite = *models.CreateRandomSprite(bounds)
		mutex.Unlock()
	}
}

func handleInput(win *pixelgl.Window) {
	for !win.Closed() {
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			pos := win.MousePosition()

			mutex.Lock()
			rect := sprite.Sprite.Frame().Moved(sprite.Matrix.Project(pixel.ZV))
			if rect.Contains(pos) && !sprite.Clicked {
				sprite.Clicked = true
				counter++
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

	sprite = *models.CreateRandomSprite(cfg.Bounds)
	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		mutex.Lock()
		if !sprite.Clicked {
			sprite.Sprite.Draw(win, sprite.Matrix)
		}

		txt.Clear()
		fmt.Fprintf(txt, "Counter: %d\nTimer: %02d:%02d", counter, timer/60, timer%60)
		txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))

		mutex.Unlock()

		win.Update()
	}
}
