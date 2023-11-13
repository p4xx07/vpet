package pet

import (
	"math/rand"
	"time"
)

type IPet interface {
	Feed()
	Play()
	GetFrames() []string
	IsEating() bool
	IsPlaying() bool
}

type Pet struct {
	Name      string
	Type      string
	Hunger    int
	Happiness int
	Strength  int
	Birthdate time.Time
	LastFed   time.Time
}

var names = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Kathy", "Leo", "Molly", "Nathan", "Olivia", "Paul", "Quincy", "Rachel", "Sam", "Tom", "Ursula", "Victor", "Wendy", "Xander", "Yvonne", "Zane"}
var petTypes = []string{"Dog", "Cat", "Bird", "Fish", "Rabbit", "Hamster", "Guinea Pig", "Ferret", "Turtle", "Snake", "Lizard", "Horse", "Pony", "Gerbil", "Mouse", "Rat", "Chinchilla", "Hedgehog", "Tarantula", "Hermit Crab"}

func NewPet() *Pet {
	return &Pet{
		Name:      names[rand.Intn(len(names))],
		Type:      petTypes[rand.Intn(len(petTypes))],
		Hunger:    0,
		Happiness: 0,
		Strength:  rand.Intn(255),
		Birthdate: time.Now(),
		LastFed:   time.Now(),
	}
}

func (p *Pet) Feed() {
	p.Hunger = 0
	p.LastFed = time.Now()
}

func (p *Pet) Play() {
	p.Happiness += 10
}

func (p *Pet) GetFrames() []string {
	return []string{
		`/ \__
(    @\___
/         O
/   (_____/
/_____/   
	`,
		`/ \__
(    @\___
/         O
/   (_____/
/_____/   U`,
		`/ \__
(    !\___
/         O
/   (_____/
/_____/   U`,
		`/ \__
(    #\___
/         🔥
/   (_____/
/_____/  U`,
	}
}
