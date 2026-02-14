package main

import (
	"math"
	"time"

	r "github.com/gen2brain/raylib-go/raylib"
)

func CalculateAlfaXY(angle, rad float32) (float32, float32, float32) {
	alfa := (angle-90) * math.Pi / 180.0 
	return alfa, rad * float32(math.Cos(float64(alfa))), rad * float32(math.Sin(float64(alfa))) 
}

func main() {
	r.InitWindow(500, 500, "Clock v0.0.1")
	r.SetTargetFPS(60)
		
	radS := 90.0
	radM := 70.0
	radH := 50.0

	Time := time.Now()
	angleSec := float64(Time.Second()) * 6.0
	angleMin := float64(Time.Minute()) * 6.0 + float64(Time.Second()) * 0.1
	angleHour := float64(Time.Hour()%12) * 30.0 + float64(Time.Minute()) * 0.5
	FPS := 0

	camera := r.Camera2D{
		Offset: r.Vector2{X: 250, Y: 250},
		Target: r.Vector2{X: 0, Y: 0},
		Rotation: 0.0,
		Zoom: 1.0,
	}
	
	for !r.WindowShouldClose() {
		FPS++
		r.BeginDrawing()
		r.ClearBackground(r.Black)
		r.BeginMode2D(camera)
		r.DrawCircle(0, 0, 100.0, r.White)
		r.DrawCircle(0, 0, 99.0, r.Black)
		
		alfaSec, xSec, ySec := CalculateAlfaXY(float32(angleSec), float32(radS))
		alfaMin, xMin, yMin := CalculateAlfaXY(float32(angleMin), float32(radM))
		alfaHour, xHour, yHour := CalculateAlfaXY(float32(angleHour), float32(radH))

		r.DrawLineEx(r.Vector2{0,0}, r.Vector2{float32(xHour), float32(yHour)}, 1.0, r.Red)
		r.DrawLineEx(r.Vector2{0,0}, r.Vector2{float32(xMin), float32(yMin)}, 1.0, r.Blue)
		r.DrawLineEx(r.Vector2{0,0}, r.Vector2{float32(xSec), float32(ySec)}, 1.0, r.Yellow)
		r.DrawCircle(0, 0, 2.0, r.White)
		if FPS == 60 {
			angleHour += 0.0083
			angleMin += 0.1
			angleSec += 6.0

			if alfaSec >= 360.0 {
				angleSec = 0.0
			} 
			if alfaMin >= 360.0 {
				angleMin = 0.0
			}
			if alfaHour >= 360.0 {
				angleHour = 0.0
			}
			FPS = 0
		}

		r.EndMode2D()
		r.EndDrawing()
	}
}
