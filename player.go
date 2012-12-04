package main

import (
  "github.com/jteeuwen/glfw"
  ry "../ridley"

)

type ControlComponent struct {
  owner *ry.Object 
}

func newControlComponent(owner *ry.Object) *ControlComponent {
  return &ControlComponent{owner}
}

func (c ControlComponent) Update() {
  o := c.owner
  if glfw.Key('E') == glfw.KeyPress {
    o.Position.Z -= 0.1;
  } else if glfw.Key('D') == glfw.KeyPress {
    o.Position.Z += 0.1;
  } else if glfw.Key('S') == glfw.KeyPress {
    o.Position.X -= 0.1;
  } else if glfw.Key('F') == glfw.KeyPress {
    o.Position.X += 0.1;
  }
}

