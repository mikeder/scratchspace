package main

import (
	_ "embed"
	"scratch/raylib/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600

	spriteScale = 8
)

type AnimatedSprite2D struct {
	Position rl.Vector2
	Rotation float32
	Scale    float32
	Size     rl.Vector2

	accumulator float32
	fps         float32
	srcRect     rl.Rectangle
	dstRect     rl.Rectangle
	origin      rl.Vector2
	texture     rl.Texture2D
}

func NewAnimatedSprite2D(fps float32, position rl.Vector2, scale float32, size rl.Vector2, texture rl.Texture2D) *AnimatedSprite2D {
	return &AnimatedSprite2D{
		Position: position,
		Rotation: 0,
		Scale:    scale,
		Size:     size,

		accumulator: 0,
		fps:         1.0 / fps,
		srcRect:     rl.NewRectangle(0, 0, size.X, size.Y),
		dstRect:     rl.NewRectangle(position.X, position.Y, size.X*scale, size.Y*scale),
		origin:      rl.NewVector2(0, 0),
		texture:     texture,
	}
}

func (s *AnimatedSprite2D) update(delta float32) {
	s.accumulator += delta

	if s.accumulator >= s.fps {
		s.srcRect.X += s.Size.X
		if s.srcRect.X >= float32(s.texture.Width) {
			s.srcRect.X = 0
		}

		s.accumulator -= s.fps
	}

	rl.DrawTexturePro(s.texture, s.srcRect, s.dstRect, s.origin, s.Rotation, rl.White)
}

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable)
	rl.SetTargetFPS(60)
	rl.InitWindow(screenWidth, screenHeight, "raylib scratch pad")
	defer rl.CloseWindow()

	gopherSheetImg := rl.LoadImageFromMemory(".png", assets.GopherSheetPNGBytes, int32(len(assets.GopherSheetPNGBytes)))
	gopherSheet := rl.LoadTextureFromImage(gopherSheetImg)
	defer rl.UnloadTexture(gopherSheet)
	rl.UnloadImage(gopherSheetImg)

	sprite := NewAnimatedSprite2D(4, rl.NewVector2(float32(screenWidth/2), float32(screenWidth/2)), spriteScale, rl.NewVector2(12, 14), gopherSheet)

	var delta float32
	lastTime := rl.GetTime()
	var update = func() {
		// calculate delta
		now := rl.GetTime()
		delta = float32(now - lastTime)
		lastTime = now

		// gather input

		// update positions
		sprite.update(delta)

		// draw entities
		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

		rl.EndDrawing()
	}

	// Needed for WASM build
	// rl.SetMainLoop(update)

	for !rl.WindowShouldClose() {
		update()
	}
}
