package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"virtualpet/instructor"
	"virtualpet/painter"
	"virtualpet/pet"
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

	p := &pet.Pet{}
	command := ""
	painter.NewUi(screen, &command, p)
	i := instructor.NewService(p)

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			continue
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEnter:
				i.Handle(command)
				command = ""
			case tcell.KeyRune:
				command += string(ev.Rune())
			}
		}
	}
}
