package pet

import (
	"math/rand"
	"strings"
	"time"
)

type IPet interface {
	GetPet() *Pet
	Feed()
	Play()
	GetFrames() []string
	IsEating() bool
	IsPlaying() bool
	GetLocation() string
	SetLocation(s string)
}

type Pet struct {
	Name      string
	Type      string
	Hunger    int
	Happiness int
	Location  string
	Birthdate time.Time
	LastFed   time.Time
}

var names = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Kathy", "Leo", "Molly", "Nathan", "Olivia", "Paul", "Quincy", "Rachel", "Sam", "Tom", "Ursula", "Victor", "Wendy", "Xander", "Yvonne", "Zane"}
var petTypes = []string{"Dog", "Cat", "Bird", "Fish", "Bunny"}

func NewPet() IPet {
	p := Pet{
		Name:      names[rand.Intn(len(names))],
		Type:      petTypes[rand.Intn(len(petTypes))],
		Hunger:    rand.Intn(255),
		Happiness: rand.Intn(255),
		Location:  "home",
		Birthdate: time.Now(),
		LastFed:   time.Now(),
	}

	return choosePet(p)
}

func choosePet(pet Pet) IPet {
	switch strings.ToLower(pet.Type) {
	case "cat":
		return &Cat{Pet: pet}
	case "bird":
		return &Bird{Pet: pet}
	case "fish":
		return &Fish{Pet: pet}
	case "bunny":
		return &Bunny{Pet: pet}
	default:
		return &Dog{Pet: pet}
	}

}

func (p *Pet) Feed() {
	p.Hunger = 0
	p.LastFed = time.Now()
}

func (p *Pet) SetLocation(location string) {
	p.Location = location
}

func (p *Pet) GetLocation() string {
	return p.Location
}

func (p *Pet) Play() {
	p.Happiness += 10
}

func (p *Pet) GetFrames() []string {
	return []string{"N/A"}
}

func (p *Pet) IsPlaying() bool {
	return true
}

func (p *Pet) IsEating() bool {
	return true
}

func (p *Pet) GetPet() *Pet {
	return p
}
