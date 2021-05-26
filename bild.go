package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Maler interface {
	Male(bildschirm *ebiten.Image)
	Reihenfolge() int
}

type Bild struct {
	bild         *ebiten.Image
	bildOptionen *ebiten.DrawImageOptions
	entfernung   int
}

func (bild *Bild) Male(bildschirm *ebiten.Image) {
	if bild.bild != nil {
		bildschirm.DrawImage(bild.bild, bild.bildOptionen)
	}
}

func (bild *Bild) Reihenfolge() int {
	return bild.entfernung
}
