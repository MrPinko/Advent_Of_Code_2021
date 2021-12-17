package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Probe struct {
	x     int32
	y     int32
	vel_x int32
	vel_y int32
}

type Target struct {
	min_x int32
	min_y int32
	max_x int32
	max_y int32
}

type Camera_Dummy struct {
	X int32
	Y int32
}

const (
	velocity int32 = 10
)

var (
	scale        int32        = 10
	probe        Probe        = Probe{x: 0, y: 0, vel_x: velx, vel_y: vely}
	target       Target       = Target{min_x: 20, max_x: 30, min_y: 5, max_y: 10}
	camera_Dummy Camera_Dummy = Camera_Dummy{0, 0}
	target_rect  rl.Rectangle = rl.Rectangle{}
	probe_rect   rl.Rectangle = rl.Rectangle{}
	hitted       bool         = false
)

var (
	velx int32 = 6
	vely int32 = 9
)

func update() {
	probe.x += probe.vel_x
	probe.y -= probe.vel_y //!go up because y is inverted

	if probe.vel_x != 0 {
		if probe.vel_x > 0 {
			probe.vel_x -= 1
		} else if probe.vel_x < 0 {
			probe.vel_x += 1
		}
	}
	probe.vel_y -= (1) //gravity
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Trench Destroyer (with style)")

	rl.SetTargetFPS(60)

	camera := rl.Camera2D{}
	camera.Target = rl.NewVector2(float32(camera_Dummy.X), float32((camera_Dummy.Y)))
	camera.Offset = rl.NewVector2(float32((screenWidth / 2)), float32((screenHeight / 2)))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	lines := []rl.Vector2{{X: 0, Y: 0}}

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) && !hitted {
			update()
			lines = append(lines, rl.Vector2{X: float32(probe.x * scale), Y: float32(probe.y * scale)})
		}

		if rl.CheckCollisionRecs(probe_rect, target_rect) {
			hitted = true
		}

		if rl.IsKeyDown(rl.KeyD) {
			camera_Dummy.X += 1 * velocity // Player movement
		} else if rl.IsKeyDown(rl.KeyA) {
			camera_Dummy.X -= 1 * velocity // Player movement
		}
		if rl.IsKeyDown(rl.KeyW) {
			camera_Dummy.Y -= 1 * velocity // Player movement
		} else if rl.IsKeyDown(rl.KeyS) {
			camera_Dummy.Y += 1 * velocity // Player movement
		}

		camera.Target = rl.NewVector2(float32(camera_Dummy.X), float32(camera_Dummy.Y))

		// Camera rotation controls
		if rl.IsKeyDown(rl.KeyQ) {
			camera.Rotation--
		} else if rl.IsKeyDown(rl.KeyE) {
			camera.Rotation++
		}

		// Limit camera rotation to 80 degrees (-40 to 40)
		if camera.Rotation > 40 {
			camera.Rotation = 40
		} else if camera.Rotation < -40 {
			camera.Rotation = -40
		}

		// Camera zoom controls
		camera.Zoom += float32(rl.GetMouseWheelMove()) * 0.05

		if camera.Zoom > 3.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom < 0.1 {
			camera.Zoom = 0.1
		}

		// Camera reset (zoom and rotation)
		if rl.IsKeyPressed(rl.KeyR) {
			camera.Zoom = 1.0
			camera.Rotation = 0.0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(camera)

		rl.ClearBackground(rl.DarkGray)

		//draw target
		target_rect = rl.NewRectangle(float32(target.min_x*scale), float32((target.min_y * scale)),
			float32((target.max_x-target.min_x+1)*scale), float32((target.max_y-target.min_y+1)*scale))
		rl.DrawRectangleRec(target_rect, rl.Color{R: 0, B: 255, G: 0, A: 255})

		//draw probe
		probe_rect = rl.NewRectangle(float32(probe.x*scale), float32(probe.y*scale), float32(1*scale), float32(1*scale))
		rl.DrawRectangleRec(probe_rect, rl.Color{R: 255, B: 0, G: 0, A: 255})

		if len(lines) >= 2 {
			for i := 0; i < len(lines)-1; i++ {
				rl.DrawLineEx(lines[i], lines[i+1], 1, rl.Blue)
			}
		}

		rl.EndMode2D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
