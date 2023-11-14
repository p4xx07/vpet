package pet

import (
	"math/rand"
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
	Strength  int
	Location  string
	Birthdate time.Time
	LastFed   time.Time
}

var names = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Kathy", "Leo", "Molly", "Nathan", "Olivia", "Paul", "Quincy", "Rachel", "Sam", "Tom", "Ursula", "Victor", "Wendy", "Xander", "Yvonne", "Zane"}
var petTypes = []string{"Dog", "Cat", "Bird", "Fish", "Rabbit", "Hamster", "Guinea Pig", "Ferret", "Turtle", "Snake", "Lizard", "Horse", "Pony", "Gerbil", "Mouse", "Rat", "Chinchilla", "Hedgehog", "Tarantula", "Hermit Crab"}

func NewPet() IPet {
	p := Pet{
		Name:      names[rand.Intn(len(names))],
		Type:      petTypes[rand.Intn(len(petTypes))],
		Hunger:    0,
		Happiness: 0,
		Location:  "Home",
		Strength:  rand.Intn(255),
		Birthdate: time.Now(),
		LastFed:   time.Now(),
	}

	return choosePet(p)
}

func choosePet(pet Pet) IPet {
	randomNumber := rand.Intn(4) + 1
	switch randomNumber {
	case 1:
		return &Cat{Pet: pet}
	case 2:
		return &Bird{Pet: pet}
	case 3:
		return &Fish{Pet: pet}
	case 4:
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
