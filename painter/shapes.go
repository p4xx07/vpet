package painter

import "github.com/gdamore/tcell/v2"

var defaultStyle = tcell.StyleDefault.
	Foreground(tcell.ColorGreen).
	Background(tcell.ColorReset)

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
