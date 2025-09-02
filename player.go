package main

import rl "github.com/gen2brain/raylib-go/raylib"

type AnimationPlayer struct {
	spriteSheet    rl.Texture2D
	animationNames []string
	animations     []Animation
}

type Animation struct {
	frameSize                       rl.Vector2
	yIndex, currentFrame, animFrame int32
}

func (a *Animation) Start() {
	a.currentFrame = 0
	a.animFrame = 0
}

func (a *Animation) Play() {

	a.animFrame++
}

func (a *Animation) Init(frameSize rl.Vector2, yIndex int32) {
	a.frameSize = frameSize
	a.yIndex = yIndex
}
