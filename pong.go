package main

import (
	"fmt"
	"os"
	"time"

	img "github.com/veandco/go-sdl2/img"

	sdl "github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	height       = 600
	width        = 800
	playerHight  = 120
	playerWidth  = 32
	playerBorder = 10
	ballHeight   = 16
	ballWidth    = 16
)

var xVelocity int32 = 3 / 2
var yVelocity int32 = 3 / 2
var hitBox1 = &sdl.Rect{X: width - playerWidth, Y: 240, W: playerWidth, H: playerHight}
var hitBox2 = &sdl.Rect{X: 0, Y: 240, W: playerWidth, H: playerHight}
var ballBox = &sdl.Rect{X: width/2 - ballWidth/2, Y: height/2 - ballHeight/2, W: ballWidth, H: ballHeight}
var p1score = 0
var p2score = 0
var scoreBox1 = &sdl.Rect{X: 450, Y: 100, W: 32, H: 32}
var scoreBox2 = &sdl.Rect{X: 350 - 32, Y: 100, W: 32, H: 32}
var numbers = []string{"resources/images/zero.png",
	"resources/images/one.png",
	"resources/images/two.png",
	"resources/images/three.png",
	"resources/images/four.png",
	"resources/images/five.png",
	"resources/images/six.png",
	"resources/images/seven.png",
	"resources/images/eight.png",
	"resources/images/nine.png"}

func initSdl() {
	// This part initialises sdl for the project
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not init %v", err)
		os.Exit(2)
	}
	defer sdl.Quit()
}

func initTtf() {
	// initialise string output with ttf package
	if err := ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not init font %v", err)
	}

}

func createWindorAndRenderer(w, h int32) (*sdl.Window, *sdl.Renderer) {
	// creating a windor and the renderer
	window, renderer, err := sdl.CreateWindowAndRenderer(w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create window %v", err)
	}
	return window, renderer
}

func openingFont(path string) *ttf.Font {
	//now we create the title
	font, err := ttf.OpenFont(path, 20)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create font %v", err)
	}
	return font
}

func welcomeScene(title string, font *ttf.Font) *sdl.Surface {
	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	surface, err := font.RenderUTF8Solid(title, color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write out title %v", err)
	}
	return surface
}

func drawBackground(image string, renderer *sdl.Renderer, texture *sdl.Texture) {
	// background here

	background, err := img.LoadTexture(renderer, image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load background %v", err)
	}
	renderer.Copy(background, nil, nil)
}

func drawPlayersAndBall(renderer *sdl.Renderer, texture *sdl.Texture) {
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
	renderer.Copy(player1, nil, hitBox1)
	renderer.Copy(player2, nil, hitBox2)
	renderer.Copy(ball, nil, ballBox)
}

func drawPoints(renderer *sdl.Renderer) {
	scoreOne, err := img.LoadTexture(renderer, numbers[p1score])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load score image %v", err)
	}
	scoreTwo, err := img.LoadTexture(renderer, numbers[p2score])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load score image %v", err)
	}
	renderer.Copy(scoreOne, nil, scoreBox1)
	renderer.Copy(scoreTwo, nil, scoreBox2)
}

func createTextureFromSurface(renderer *sdl.Renderer, surface *sdl.Surface) *sdl.Texture {
	// creating texture from the surface
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create texture %v", err)
	}
	return texture
}

func drawTitle(renderer *sdl.Renderer, texture *sdl.Texture) {
	renderer.Clear()
	renderer.Copy(texture, nil, nil)
	renderer.Present()
	//just to see the window, as the loop comes in it will be removed
	time.Sleep(time.Second * 3)
}

func drawGame(renderer *sdl.Renderer, texture *sdl.Texture) {
	renderer.Clear()
	drawBackground("resources/images/background.png", renderer, texture)
	drawPoints(renderer)
	drawPlayersAndBall(renderer, texture)
	renderer.Present()
}

func updatePlayerOne(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_UP {
		if hitBox1.Y >= 5 {
			hitBox1.Y -= 12
		}
	} else if key.Keysym.Scancode == sdl.SCANCODE_DOWN {
		if hitBox1.Y <= height-playerHight-5 {
			hitBox1.Y += 10
		}
	}
}

func updatePlayerTwo(key *sdl.KeyboardEvent) {
	if key.Keysym.Scancode == sdl.SCANCODE_W {
		if hitBox2.Y >= 5 {
			hitBox2.Y -= 12
		}
	} else if key.Keysym.Scancode == sdl.SCANCODE_S {
		if hitBox2.Y <= height-playerHight-5 {
			hitBox2.Y += 10
		}
	}
}

func collisionDetection() {
	if ballBox.X+ballBox.W == hitBox1.X+playerBorder && ballBox.Y+ballBox.H >= hitBox1.Y && ballBox.Y < hitBox1.Y+hitBox1.H ||
		ballBox.X == hitBox2.X+hitBox2.W-playerBorder && ballBox.Y+ballBox.H >= hitBox2.Y && ballBox.Y < hitBox2.Y+hitBox2.H {
		xVelocity = -xVelocity
	}
}

func bounceFromWall() {
	if ballBox.Y < 0 || ballBox.Y > height-ballBox.H/2 {
		yVelocity = -yVelocity
	}
}

func resetBallPosition() {
	if ballBox.X < 0 {
		ballBox.X = width/2 - ballWidth/2
		ballBox.Y = height/2 - ballHeight/2
		if p1score < 9 {
			p1score++
		} else {
			p2score = 0
			p1score = 0
		}
		time.Sleep(time.Second)
	} else if ballBox.X > width {
		ballBox.X = width/2 - ballWidth/2
		ballBox.Y = height/2 - ballHeight/2
		if p2score < 9 {
			p2score++
		} else {
			p1score = 0
			p2score = 0
		}
		time.Sleep(time.Second)
	}
}

func updateBall() {

	ballBox.X += xVelocity
	ballBox.Y += yVelocity

	bounceFromWall()
	resetBallPosition()
	collisionDetection()
}
func main() {

	initSdl()
	initTtf()
	window, renderer := createWindorAndRenderer(width, height)
	font := openingFont("resources/fonts/Arial.ttf")
	surface := welcomeScene("PongGo", font)
	texture := createTextureFromSurface(renderer, surface)
	defer ttf.Quit()
	defer window.Destroy()
	defer font.Close()
	defer surface.Free()
	defer texture.Destroy()
	var running = true

	drawTitle(renderer, texture)
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch key := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				updatePlayerOne(key)
				updatePlayerTwo(key)
			}
		}
		updateBall()
		drawGame(renderer, texture)
	}
}
