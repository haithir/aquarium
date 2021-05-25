package main

import (
	"github.com/hajimehoshi/ebiten"
	. "time"
)

type AbstrakteFigur interface {
	// return: bool=true zerstöre, array neuer Figuren oder nil
	Bewege(vergangen Duration) (bool, []AbstrakteFigur)
	AlleMaler() []Maler
}

type BewegteFigur struct {
	bild Bild
	verhalten Verhalten
}

func NeueFigur(bild *ebiten.Image) BewegteFigur {
	return BewegteFigur{
		bild: Bild{
			bild:         bild,
			bildOptionen: nil,
			entfernung:   0,
		},
		verhalten: &ImmerGeradeaus{
			x: 1000,
			y: 100,
		},
	}
}

func (figur *BewegteFigur) Bewege(vergangen Duration) (bool, []AbstrakteFigur) {
	return figur.verhalten.Bewege(figur, vergangen)
}

func (figur *BewegteFigur) AlleMaler() []Maler {
	return  []Maler{&figur.bild}
}
