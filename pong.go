package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {

	// This part initialises sdl for the project
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not init %v", err)
		os.Exit(2)
	}
	defer sdl.Quit()

	// initialise string output with ttf package
	if err := ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not init font %v", err)
	}
	defer ttf.Quit()

	// creating a windor and the renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create window %v", err)
	}
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
	renderer.Copy(texture, nil, nil)
	renderer.Present()
	time.Sleep(time.Second * 5)
}
