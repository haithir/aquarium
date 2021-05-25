package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
	"log"
	"sort"
	"time"
)

func init() {
	for _,name := range bildNamen {
		img, _, err := ebitenutil.NewImageFromFile(name)
		if err != nil {
			log.Fatal(err)
		}
		geladeneBuilder[name] = img
	}
}

type Aquarium struct{
	figuren []BildElement
	zeitLetztesBild time.Time
}

func (aquarium *Aquarium) Update() error {
	vergangen := time.Since(aquarium.zeitLetztesBild)
	for _, bildElement := range aquarium.figuren {
		bildElement.ZeitVergangen(vergangen)
	}
	aquarium.zeitLetztesBild = aquarium.zeitLetztesBild.Add(vergangen)
	return nil
}

func (aquarium *Aquarium) Draw(bildschirm *ebiten.Image) {
	malerListe := []Maler{}
	for _, bildElement := range aquarium.figuren {
		malerListe = append(malerListe, bildElement.AlleMaler()...)
	}
	sort.SliceStable(malerListe, func(i, j int) bool {
		return malerListe[i].Reihenfolge() < malerListe[j].Reihenfolge()
	})
	for _, maler := range malerListe {
		maler.Male(bildschirm)
	}
}

func (g *Aquarium) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return breite, höhe
}

func main() {
	ebiten.SetWindowSize(breite, höhe)
	ebiten.SetWindowTitle("Hello, Yuuka and Aika!")
	fisch1 := NeueFigur(geladeneBuilder["yuukafisch.png"])
	fisch2 := Figur{
		bild:      Bild{
			bild:         geladeneBuilder["yuukas_geburtstagsfisch.png"],
			entfernung:   50,
		},
		verhalten: &ImmerGeradeaus{
			x:          0,
			y:          110,
			rechtslang: true,
		},
	}
	frameCounter := FrameCounter{}
	aquarium := Aquarium{
		figuren: []BildElement{
			&fisch1,
			&fisch2,
			&frameCounter,
		},
		zeitLetztesBild: time.Now(),
	}
	if err := ebiten.RunGame(&aquarium); err != nil {
		log.Fatal(err)
	}
}