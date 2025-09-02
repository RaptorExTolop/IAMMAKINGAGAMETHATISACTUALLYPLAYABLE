package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	bkgColour                 rl.Color = rl.SkyBlue
	running                   bool     = true
	screenWidth, screenHeight int32    = 1280, 720
	screenTitle               string   = "We make a game"

	// background
	backgroundRenderer Background
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, screenTitle)
	rl.SetTargetFPS(60)

	// background inits
	backgroundRenderer.AddLayer("assets/world/background/background_layer_1.png", 320, 180, float32(screenWidth), float32(screenHeight))
	backgroundRenderer.AddLayer("assets/world/background/background_layer_2.png", 320, 180, float32(screenWidth), float32(screenHeight))
	backgroundRenderer.AddLayer("assets/world/background/background_layer_3.png", 320, 180, float32(screenWidth), float32(screenHeight))
}

func quit() {
	rl.CloseWindow()
	backgroundRenderer.Unload()
}

func main() {
	defer quit()
	for running {
		input()
		update()
		draw()
	}
}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()
}

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColour)
	backgroundRenderer.Draw()

	rl.EndDrawing()
}
