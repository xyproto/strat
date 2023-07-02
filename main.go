package main

import (
	"log"
	"path/filepath"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalln("Init Error:", err)
	}
	// Make sure to quit when the function returns
	defer sdl.Quit()

	// Create the window
	win, err := sdl.CreateWindow("Hello World!", 100, 100, 320, 200, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalln("CreateWindow Error:", err)
	}
	defer win.Destroy()

	// Create a renderer
	ren, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalln("CreateRenderer Error:", err)
	}
	defer ren.Destroy()

	// Load the image as a texture
	spaceShipsTexture, err := img.LoadTexture(ren, filepath.Join(".", "img", "spaceships.png"))
	if err != nil {
		log.Fatalln("LoadTexture Error:", err)
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "LoadTexture: %s\n", err)
	}
	defer spaceShipsTexture.Destroy()

	// No need for the image data after the texture has been created
	//png.Free()

	for i := 0; i < 40; i++ {
		// Clear the renderer and display the image/texture
		ren.Clear()
		ren.Copy(spaceShipsTexture, nil, nil)
		ren.Present()

		// Wait 100 ms
		sdl.Delay(100)
	}
}
