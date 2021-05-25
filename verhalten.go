package main

import (
	"github.com/hajimehoshi/ebiten"
	. "time"
)

type Verhalten interface {
	Bewege(figur *BewegteFigur, vergangen Duration) (bool, []AbstrakteFigur)
}

type ImmerGeradeaus struct {
	x int
	y int
	rechtslang bool
}
func (verhalten *ImmerGeradeaus) Bewege(figur *BewegteFigur, vergangen Duration) (bool, []AbstrakteFigur) {
	if verhalten.rechtslang {
		verhalten.x++
		width, _ := figur.bild.bild.Size()
		if verhalten.x > breite {
			verhalten.x = -width
		}
	} else {
		verhalten.x--
		width, _ := figur.bild.bild.Size()
		if verhalten.x < -width {
			verhalten.x = breite
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(verhalten.x), float64(verhalten.y))
	figur.bild.bildOptionen = op
	return false, nil
}

