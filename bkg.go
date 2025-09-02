package main

import rl "github.com/gen2brain/raylib-go/raylib"

type BkgLayer struct {
	Texture   rl.Texture2D
	Src, Dest rl.Rectangle
	Origin    rl.Vector2
	Rotation  float32
}

func (b *BkgLayer) Init(texturePath string, imgWidth, imgHeight, screenWidth, screenHeight float32) {
	b.Texture = rl.LoadTexture(texturePath)
	b.Src = rl.Rectangle{0, 0, imgWidth, imgHeight}
	b.Dest = rl.Rectangle{0, 0, screenWidth, screenHeight}
	b.Origin = rl.Vector2{0, 0}
	b.Rotation = 0
}

func (b *BkgLayer) Draw() {
	rl.DrawTexturePro(b.Texture, b.Src, b.Dest, b.Origin, b.Rotation, rl.RayWhite)
}

type Background struct {
	Layers []BkgLayer
}

func (b *Background) AddLayer(texturePath string, imgWidth, imgHeight, screenWidth, screenHeight float32) {
	var newBkgLayer BkgLayer
	newBkgLayer.Init(texturePath, imgWidth, imgHeight, screenWidth, screenHeight)

	b.Layers = append(b.Layers, newBkgLayer)
}

func (b *Background) Draw() {
	for _, Layer := range b.Layers {
		Layer.Draw()
	}
}

func (b *Background) Unload() {
	for _, Layer := range b.Layers {
		rl.UnloadTexture(Layer.Texture)
	}
}
