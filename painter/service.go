package painter

import (
	"fmt"
	"math/rand"
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
	go u.updateGame()
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

func (u *uiService) updateGame() {
	c := time.NewTicker(time.Millisecond * 200)

	for {
		select {
		case <-c.C:
			if rand.Intn(100) < 5 {
				u.pet.Starve(rand.Int31n(10))
			}
			if rand.Intn(100) < 5 {
				u.pet.HappinessDecay(rand.Int31n(10))
			}
			if rand.Intn(100) < 1 {
				u.pet.Shit()
			}
			u.pet.Die()
		default:
		}
	}
}

func (u *uiService) drawGame() {
	if u.pet.GetPet().Dead {
		u.drawGameOver()
		return
	}
	u.drawHeader(20)
	u.drawStats(0, 0, 20, 20)
	u.drawPet(22, 2)
	u.drawShit(44, 2)
}

func (u *uiService) drawStats(x, y, width, height int) {
	getPet := u.pet.GetPet()
	u.drawBox(x, y, width, height)
	u.drawText(x+1, y+1, fmt.Sprintf("%s | %s", getPet.Name, getPet.Type))
	u.drawText(x+1, y+2, fmt.Sprintf("📍 %s", getPet.Location))
	u.drawText(x+1, y+3, "🍕")
	u.drawProgressBar(x+5, y+3, getPet.Satisfaction, pet.MaxStat)
	u.drawText(x+1, y+4, "🎭")
	u.drawProgressBar(x+5, y+4, getPet.Happiness, pet.MaxStat)
	u.drawText(x+1, y+5, "💩")
	u.drawProgressBar(x+5, y+5, pet.MaxStat-getPet.ShitView, pet.MaxStat)
}

func (u *uiService) drawPet(x, y int) {
	animationFrames = u.pet.GetFrames()
	if u.pet == nil {
		return
	}
	u.drawShape(x, y, []rune(animationFrames[animationIndex]))
	if u.frameCounter%60 == 0 {
		animationIndex = (animationIndex + 1) % len(animationFrames)
	}
}

func (u *uiService) drawShape(x int, y int, shape []rune) {
	startX := x
	for _, t := range shape {
		if t == '\n' {
			x = startX
			y += 1
		}
		u.screen.SetContent(x, y, t, nil, defaultStyle)
		x += 1
	}
}

func (u *uiService) drawHeader(x int) {
	const header = "action: "
	for i, t := range []rune(header + *u.command) {
		u.screen.SetContent(i+x, 0, t, nil, defaultStyle)
	}
}

func (u *uiService) drawGameOver() {
	u.drawText(0, 0, fmt.Sprintf("%s is dead!", u.pet.GetPet().Name))
}

func (u *uiService) drawShit(x, y int) {
	const shitASCII = `     (   )
  (   ) (
   ) _   )
    ( \_
  _(_\ \)__
 (____\___))   
	`

	if u.pet.GetPet().ShitView <= 0 {
		return
	}
	u.drawShape(x, y, []rune(shitASCII))
}
