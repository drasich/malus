package main

import (
	"fmt"
  gl "github.com/chsc/gogl/gl21"
  "github.com/jteeuwen/glfw"
  //"github.com/drasich/ridley"
  ry "../ridley"
	"os"
  "time"

)

const (
	Title  = "malus"
	Width  = 640/3
	Height = 480/3
)

var (
  exit = false
  last_time time.Time
  //scene Scene
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, gl.TRUE)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 32, 32, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

  var scene ry.Scene
	if err := scene.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "init: %s\n", err)
		return
	}
	defer scene.Destroy()

  var player,wall ry.Object
  wall.Init()
  player.Init()
  var mat ry.Matrix4
  mat.Translation(0,0,-7)
  mat.Rotate(-90, 1,0,0)
  player.Matrix = mat
  cc := newControlComponent(&player)
  mesh := ry.NewMeshComponent("model/tex.bin", &player)
  mesh.Init()
  player.AddComponent(*cc)
  player.Mesh = mesh

  mesh2 := ry.NewMeshComponent("model/tex.bin", &wall)
  mesh2.Init()
  wall.Mesh = mesh2
  wall.Matrix = mat

  scene.AddObject(&player)
  scene.AddObject(&wall)

  last_time = time.Now()
	for glfw.WindowParam(glfw.Opened) == 1 && !exit {
    scene.Update()
		scene.Draw()
		glfw.SwapBuffers()
    since := time.Since(last_time).Seconds()
    if (since > 0.02) {
      fmt.Println("frame under 50fps:", since)
    }
    last_time = time.Now()
  }
}

