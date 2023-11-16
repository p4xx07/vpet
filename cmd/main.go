package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"math/rand"
	"os"
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

	p := pet.NewPet()
	command := ""
	painter.NewUi(screen, &command, p)

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			continue
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyDEL:
				if len(command) == 0 {
					continue
				}

				command = command[:len(command)-1]
			case tcell.KeyEnter:
				handle(screen, p, command)
				command = ""
			case tcell.KeyRune:
				command += string(ev.Rune())
			}
		}
	}
}

func handle(screen tcell.Screen, pet pet.IPet, c string) {
	switch c {
	case "play":
		pet.Play()
	case "feed":
		pet.Feed(rand.Int31n(30))
	case "clean":
		pet.Clean()
	case "park":
		pet.SetLocation("park")
	case "home":
		pet.SetLocation("home")
	case "exit":
		screen.Clear()
		screen.Show()
		os.Exit(0)
	}
}
