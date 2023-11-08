package painter

import (
	"github.com/gdamore/tcell/v2"
	"virtualpet/pet"
)

type IUI interface {
	Draw()
}

type ui struct {
	drawer  IPainter
	pet     *pet.Pet
	command *string
}

func NewUi(screen tcell.Screen, pet *pet.Pet) IUI {
	u := &ui{drawer: NewDrawer(screen), pet: pet}
	go u.Draw()
	return u
}

func (u *ui) Draw() {
	u.drawer.AddFunctionToDrawLoop(func() {
		u.drawHeader(u.drawer)
		u.drawPet()
		u.drawStats()
	})
}

func (u *ui) drawHeader(d IPainter) {
	const header = "action: "
	for i, t := range []rune(header + *u.command) {
		d.SetContent(i, 0, t, nil, defaultStyle)
	}
}

var animationFrames []string
var animationIndex int

func (u *ui) drawPet() {
	animationFrames = u.pet.GetFrames()
	if u.pet == nil {
		return
	}
	x := 0
	y := 2
	for _, t := range []rune(animationFrames[animationIndex]) {
		if t == '\n' {
			x = 0
			y += 1
		}
		u.drawer.SetContent(x, y, t, nil, defaultStyle)
		x += 1
	}
	animationIndex = (animationIndex + 1) % len(animationFrames)
}

func (u *ui) drawStats() {

}
