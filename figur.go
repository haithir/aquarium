package main

import (
	. "time"
)

type BildElement interface {
	// return: bool=true zerstöre, array neuer Figuren oder nil
	ZeitVergangen(vergangen Duration) (bool, []BildElement)
	AlleMaler() []Maler
}

type Figur struct {
	bild      Bild
	verhalten Verhalten
}

func (figur *Figur) ZeitVergangen(vergangen Duration) (bool, []BildElement) {
	return figur.verhalten.ZeitVergangen(figur, vergangen)
}

func (figur *Figur) AlleMaler() []Maler {
	return []Maler{&figur.bild}
}
