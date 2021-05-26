package main

import "github.com/hajimehoshi/ebiten/v2"

const breite = 2000
const höhe = 1000

var bildNamen = []string{
	"yuukafisch.png",
	"yuukas_geburtstagsfisch.png",
}

var geladeneBuilder = map[string]*ebiten.Image{}
