package main

import (
	"github.com/hajimehoshi/ebiten"
	. "time"
)

type BildElement interface {
	// return: bool=true zerstöre, array neuer Figuren oder nil
	ZeitVergangen(vergangen Duration) (bool, []BildElement)
	AlleMaler() []Maler
}

type Figur struct {
	bild Bild
	verhalten Verhalten
}

func NeueFigur(bild *ebiten.Image) Figur {
	return Figur{
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

func (figur *Figur) ZeitVergangen(vergangen Duration) (bool, []BildElement) {
	return figur.verhalten.ZeitVergangen(figur, vergangen)
}

func (figur *Figur) AlleMaler() []Maler {
	return  []Maler{&figur.bild}
}
