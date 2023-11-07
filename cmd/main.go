package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("can't create screen", err)
	}
	if err = screen.Init(); err != nil {
		log.Fatal("can't initialize screen", err)
	}
	defer screen.Fini()

	NewUi(screen, &Pet{})

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			continue
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEnter:
				handle(command)
			case tcell.KeyRune:
				command += string(ev.Rune())
			}
		}
	}
}
