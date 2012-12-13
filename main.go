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
  player.Position.Z = -7
  player.Position.X = 3
  player.Orientation = ry.QuatAngleAxis(-45*ry.DegToRad, ry.Vec3{0,0,1})

  cc := newControlComponent(&player)
  mesh := ry.NewMeshComponent("model/tex.bin", &player)
  mesh.Init()
  player.AddComponent(*cc)
  player.Mesh = mesh

  mesh2 := ry.NewMeshComponent("model/tex.bin", &wall)
  mesh2.Init()
  wall.Mesh = mesh2
  wall.Position.Z = -10
  box := ry.NewBoxComponent(ry.Vec3{0,0,0},ry.Vec3{10,10,10})
  wall.Box = box
  wall.Orientation = ry.QuatAngleAxis(-90*ry.DegToRad, ry.Vec3{0,0,1})

  scene.AddObject(&player)
  scene.AddObject(&wall)

  //plane := ry.Plane{ry.Vec3{0,0,0}, ry.Vec3{1,1,0}}
  ray := ry.Ray{ry.Vec3{15,-5,-5}, ry.Vec3{-100,0,0}}

/*
  b, hv := ry.IntersectionRayPlane(ray, plane)
  if b {
    fmt.Println("hv : ", hv)
  } else { fmt.Println("no collision : ", hv) }
  */

  //aabox := ry.AABox{ry.Vec3{0,0,0}, ry.Vec3{10,10,10}}
  //hit, _, pos, nor := ry.IntersectionRayAABox(ray, aabox)
  hit, _, pos, nor := ry.IntersectionRayObject(ray, &wall)
  if hit {
    fmt.Println("pos and normal : ", pos, nor)
  } else {
    fmt.Println("no intersection with box")

  }

  last_time = time.Now()
	for glfw.WindowParam(glfw.Opened) == 1 && !exit {
  /*
    objects, positions := ry.LaunchRay( ry.Vec3{0,0,0}, ry.Vec3{0,0,-1}, 100, scene.Objects) 

     if objects != nil {
       fmt.Println("collision with one or more objects", positions[0])
     }
     */

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

