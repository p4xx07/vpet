package main

import (
	"github.com/gdamore/tcell/v2"
)

type IUI interface {
	Draw()
}

type ui struct {
	drawer IPainter
	pet    *Pet
}

func NewUi(screen tcell.Screen, pet *Pet) IUI {
	u := &ui{drawer: NewDrawer(screen), pet: pet}
	go u.Draw()
	return u
}

func (u *ui) Draw() {
	u.drawer.AddFunctionToDrawLoop(func() {
		u.drawHeader(u.drawer)
		u.drawPet(u.drawer, u.pet)
	})
}

func (u *ui) drawHeader(d IPainter) {
	const header = "action: "
	for i, t := range []rune(header + command) {
		d.SetContent(i, 0, t, nil, defaultStyle)
	}
}

var animationFrames []string
var animationIndex int

func (u *ui) drawPet(drawer IPainter, pet *Pet) {
	animationFrames = pet.GetFrames()
	if pet == nil {
		return
	}
	x := 0
	y := 2
	for _, t := range []rune(animationFrames[animationIndex]) {
		if t == '\n' {
			x = 0
			y += 1
		}
		drawer.SetContent(x, y, t, nil, defaultStyle)
		x += 1
	}
	animationIndex = (animationIndex + 1) % len(animationFrames)
}
