package painter

import (
	"fmt"
	"sync"
	"time"
	"virtualpet/pet"

	"github.com/gdamore/tcell/v2"
)

type IUIService interface{}

type uiService struct {
	pet           pet.IPet
	command       *string
	screen        tcell.Screen
	mutex         sync.RWMutex
	functions     []func()
	functionMutex sync.RWMutex
	frameCounter  int
}

func NewUi(screen tcell.Screen, command *string, pet pet.IPet) IUIService {
	u := &uiService{pet: pet, command: command, screen: screen}
	go u.Start()
	return u
}

func (u *uiService) Start() {
	for {
		u.screen.Clear()
		u.functionMutex.RLock()

		u.drawGame()

		u.functionMutex.RUnlock()
		u.screen.Show()

		u.frameCounter++
		time.Sleep(time.Millisecond * 15)
	}
}

func (u *uiService) drawGame() {
	u.drawHeader(20)
	u.drawStats(0, 0, 20, 20)
	u.drawPet(22, 2)
}

func (u *uiService) drawStats(x, y, width, height int) {
	getPet := u.pet.GetPet()
	u.drawBox(x, y, width, height)
	u.drawText(x+1, y+1, fmt.Sprintf("Name: %s", getPet.Name))
	u.drawText(x+1, y+2, fmt.Sprintf("Type: %s", getPet.Type))
	u.drawText(x+1, y+3, fmt.Sprintf("Hunger: %d", getPet.Hunger))
	u.drawText(x+1, y+4, fmt.Sprintf("Happiness: %d", getPet.Happiness))
	u.drawText(x+1, y+5, fmt.Sprintf("Location: %s", getPet.Location))
}

func (u *uiService) drawPet(x, y int) {
	startX := x
	animationFrames = u.pet.GetFrames()
	if u.pet == nil {
		return
	}
	for _, t := range []rune(animationFrames[animationIndex]) {
		if t == '\n' {
			x = startX
			y += 1
		}
		u.screen.SetContent(x, y, t, nil, defaultStyle)
		x += 1
	}
	if u.frameCounter%60 == 0 {
		animationIndex = (animationIndex + 1) % len(animationFrames)
	}
}

func (u *uiService) drawHeader(x int) {
	const header = "action: "
	for i, t := range []rune(header + *u.command) {
		u.screen.SetContent(i+x, 0, t, nil, defaultStyle)
	}
}
