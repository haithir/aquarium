package main

import (
	"github.com/hajimehoshi/ebiten"
	. "time"
)

type Verhalten interface {
	ZeitVergangen(figur *Figur, vergangen Duration) (bool, []BildElement)
}

type ImmerGeradeaus struct {
	x int
	y int
	rechtslang bool
}
func (verhalten *ImmerGeradeaus) ZeitVergangen(figur *Figur, vergangen Duration) (bool, []BildElement) {
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

