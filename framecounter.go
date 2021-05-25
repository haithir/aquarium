package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	. "time"
)

type FrameCounter struct {
	gezählteZeitDauer int64
	frames int
	framesPerSecond int
}

func (c *FrameCounter) Male(bildschirm *ebiten.Image)  {
	ebitenutil.DebugPrintAt(bildschirm, fmt.Sprint("frames per second: ", c.framesPerSecond), 30, 200)
}

func (*FrameCounter) Reihenfolge() int {
	return 1000000
}

func (c *FrameCounter) ZeitVergangen(vergangen Duration) (bool, []BildElement) {
	c.gezählteZeitDauer += vergangen.Milliseconds()
	c.frames++
	if c.gezählteZeitDauer >= 1000 { // 1s
		c.framesPerSecond = c.frames
		c.gezählteZeitDauer-= 1000
		c.frames = 0
	}
	return false, nil
}

func (c *FrameCounter) AlleMaler() []Maler {
	return []Maler{c}
}