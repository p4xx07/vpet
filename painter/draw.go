package painter

import (
	"github.com/gdamore/tcell/v2"
	"sync"
	"time"
)

type IPainter interface {
	Start()
	AddFunctionToDrawLoop(f func())
	SetContent(x, y int, primary rune, combining []rune, style tcell.Style)
}

type painter struct {
	screen        tcell.Screen
	mutex         sync.RWMutex
	functions     []func()
	functionMutex sync.RWMutex
}

var defaultStyle = tcell.StyleDefault.
	Foreground(tcell.ColorGreen).
	Background(tcell.ColorReset)

func NewDrawer(screen tcell.Screen) IPainter {
	d := painter{screen: screen, functions: make([]func(), 0)}
	d.Start()
	return &d
}

func (d *painter) Start() {
	go func() {
		for {
			d.screen.Clear()
			d.functionMutex.RLock()
			for _, f := range d.functions {
				f()
			}
			d.functionMutex.RUnlock()
			d.screen.Show()
			time.Sleep(time.Millisecond * 15)
		}
	}()
}

func (d *painter) SetContent(x, y int, primary rune, combining []rune, style tcell.Style) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.screen.SetContent(x, y, primary, combining, style)
}

func (d *painter) AddFunctionToDrawLoop(f func()) {
	d.functions = append(d.functions, f)
}

func (d *painter) ClearDrawLoop() {
	d.functionMutex.Lock()
	defer d.functionMutex.Unlock()
	d.functions = make([]func(), 0)
}
