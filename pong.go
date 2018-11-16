package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func initSdl() {
	// This part initialises sdl for the project
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not init %v", err)
		os.Exit(2)
	}

}

func initTtf() {
	// initialise string output with ttf package
	if err := ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not init font %v", err)
	}
}

func createWindorAndRenderer() (*sdl.Window, *sdl.Renderer) {
	// creating a windor and the renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create window %v", err)
	}
	return window, renderer
}

func main() {

	initSdl()
	initTtf()
	window, renderer := createWindorAndRenderer()
	defer sdl.Quit()
	defer ttf.Quit()
	defer window.Destroy()
	renderer.Clear()
	//now we create the title
	font, err := ttf.OpenFont("resources/fonts/Arial.ttf", 20)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create font %v", err)
	}
	defer font.Close()

	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	surface, err := font.RenderUTF8Solid("PongGo", color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write out title %v", err)
	}
	defer surface.Free()
	// creating texture from the surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create texture %v", err)
	}
	defer texture.Destroy()
	err = renderer.Copy(texture, nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not copy texture %v", err)
	}
	renderer.Present()
	//just to see the window, as the loop comes in it will be removed
	time.Sleep(time.Second * 3)

	// background here
	renderer.Clear()
	texture, err = img.LoadTexture(renderer, "resources/images/background.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load background %v", err)
	}

	//rendering players
	hitBox1 := &sdl.Rect{X: 5, Y: 240, W: 32, H: 120}
	hitBox2 := &sdl.Rect{X: 800 - 37, Y: 240, W: 32, H: 120}
	ballBox := &sdl.Rect{X: 400 - 8, Y: 300 - 8, W: 16, H: 16}

	player1, err := img.LoadTexture(renderer, "resources/images/player.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load player1 image %v", err)
	}
	player2, err := img.LoadTexture(renderer, "resources/images/player2.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load player2 image %v", err)
	}
	ball, err := img.LoadTexture(renderer, "resources/images/ball.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load ball image %v", err)
	}

	renderer.Copy(texture, nil, nil)
	renderer.Copy(player1, nil, hitBox1)
	renderer.Copy(player2, nil, hitBox2)
	renderer.Copy(ball, nil, ballBox)
	renderer.Present()

	time.Sleep(time.Second * 5)
}
