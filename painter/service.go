package painter

import (
	"fmt"
	"sync"
	"time"
	"virtualpet/pet"

	"github.com/gdamore/tcell/v2"
)

type IUIService interface {
	Draw()
}

type uiService struct {
	pet           *pet.Pet
	command       *string
	screen        tcell.Screen
	mutex         sync.RWMutex
	functions     []func()
	functionMutex sync.RWMutex
}

func NewUi(screen tcell.Screen, command *string, pet *pet.Pet) IUIService {
	u := &uiService{pet: pet, command: command, screen: screen}
	u.Draw()
	go u.Start()
	return u
}

func (u *uiService) Draw() {
	u.AddFunctionToDrawLoop(func() {
		u.drawHeader(15)
		u.drawPet(17, 2)
		u.drawStats(0, 0, 15, 20)
	})
}

func (u *uiService) drawHeader(x int) {
	const header = "action: "
	for i, t := range []rune(header + *u.command) {
		u.screen.SetContent(i+x, 0, t, nil, defaultStyle)
	}
}

func (u *uiService) drawText(x, y int, text string) {
	for i, ch := range text {
		u.screen.SetContent(x+i, y, ch, nil, tcell.StyleDefault)
	}
}

func (u *uiService) drawTextWithBorder(x, y int, text string) {
	u.drawBox(x-1, y-1, len(text)+2, 3)

	for i, ch := range text {
		u.screen.SetContent(x+i, y, ch, nil, tcell.StyleDefault)
	}
}

func (u *uiService) drawBox(x, y, width, height int) {
	for i := 0; i < width; i++ {
		u.screen.SetContent(x+i, y, '─', nil, tcell.StyleDefault)
		u.screen.SetContent(x+i, y+height-1, '─', nil, tcell.StyleDefault)
	}

	for i := 0; i < height; i++ {
		u.screen.SetContent(x, y+i, '│', nil, tcell.StyleDefault)
		u.screen.SetContent(x+width-1, y+i, '│', nil, tcell.StyleDefault)
	}

	u.screen.SetContent(x, y, '┌', nil, tcell.StyleDefault)
	u.screen.SetContent(x+width-1, y, '┐', nil, tcell.StyleDefault)
	u.screen.SetContent(x, y+height-1, '└', nil, tcell.StyleDefault)
	u.screen.SetContent(x+width-1, y+height-1, '┘', nil, tcell.StyleDefault)
}

var animationFrames []string
var animationIndex int
var frameCounter uint64

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
	if frameCounter%60 == 0 {
		animationIndex = (animationIndex + 1) % len(animationFrames)
	}
}

var defaultStyle = tcell.StyleDefault.
	Foreground(tcell.ColorGreen).
	Background(tcell.ColorReset)

func (u *uiService) Start() {
	for {
		u.screen.Clear()
		u.functionMutex.RLock()
		for _, f := range u.functions {
			f()
		}
		u.functionMutex.RUnlock()
		u.screen.Show()

		frameCounter++
		time.Sleep(time.Millisecond * 15)
	}
}

func (u *uiService) Screen() tcell.Screen {
	return u.screen
}

func (u *uiService) AddFunctionToDrawLoop(f func()) {
	u.functions = append(u.functions, f)
}

func (u *uiService) ClearDrawLoop() {
	u.functionMutex.Lock()
	defer u.functionMutex.Unlock()
	u.functions = make([]func(), 0)
}

func (u *uiService) drawStats(x, y, width, height int) {
	u.drawBox(x, y, width, height)
	u.drawText(x+1, y+1, fmt.Sprintf("Hunger: %d", u.pet.Hunger))
	u.drawText(x+1, y+2, fmt.Sprintf("Happiness: %d", u.pet.Happiness))
	u.drawText(x+1, y+3, fmt.Sprintf("Strength: %d", u.pet.Strength))
}
