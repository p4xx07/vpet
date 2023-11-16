package pet

import (
	"math"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

const (
	MaxStat int32 = 255
	MinStat int32 = 0

	WasteCoefficient float64 = 0.8
	MaxShitLength    int32   = 50
)

type IPetActions interface {
	Feed(amount int32)
	Clean()
	Play()
}

type IPetDecay interface {
	HappinessDecay(amount int32)
	Starve(amount int32)
	Shit()
	Die()
}

type IPetChecks interface {
	GetPet() Pet
	IsEating() bool
	IsPlaying() bool
	SetLocation(s string)
	GetFrames() []string
}

type IPet interface {
	IPetActions
	IPetChecks
	IPetDecay
}

type Pet struct {
	Name         string
	Type         string
	Satisfaction int32
	Happiness    int32
	ShitBuffer   int32
	ShitView     int32
	Location     string
	Birthdate    time.Time
	LastFed      time.Time
	Dead         bool
}

var names = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Kathy", "Leo", "Molly", "Nathan", "Olivia", "Paul", "Quincy", "Rachel", "Sam", "Tom", "Ursula", "Victor", "Wendy", "Xander", "Yvonne", "Zane"}
var petTypes = []string{"Dog", "Cat", "Bird", "Fish", "Bunny"}

func NewPet() IPet {
	p := Pet{
		Name:         names[rand.Intn(len(names))],
		Type:         petTypes[rand.Intn(len(petTypes))],
		Satisfaction: rand.Int31n(MaxStat),
		Happiness:    rand.Int31n(MaxStat),
		Location:     "home",
		Birthdate:    time.Now(),
		LastFed:      time.Now(),
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

func (p *Pet) Feed(amount int32) {
	result := min(p.Satisfaction+amount, MaxStat)
	atomic.StoreInt32(&p.Satisfaction, result)
	p.LastFed = time.Now()
}

func (p *Pet) SetLocation(location string) {
	p.Location = location
}

func (p *Pet) Play() {
	const delta int32 = 10
	if p.Happiness+delta > MaxStat {
		p.Happiness = MaxStat
		return
	}

	p.Happiness += delta
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

func (p *Pet) GetPet() Pet {
	return *p
}

func (p *Pet) Shit() {
	if p.ShitBuffer == 0 {
		return
	}

	amount := min(p.ShitBuffer, rand.Int31n(MaxShitLength))
	p.ShitView += amount
	p.ShitBuffer -= amount
}

func (p *Pet) Starve(amount int32) {
	result := max(p.Satisfaction-amount, MinStat)
	delta := p.Satisfaction - result
	waste := math.Abs(float64(delta) * WasteCoefficient)
	p.ShitBuffer += int32(waste)

	atomic.StoreInt32(&p.Satisfaction, result)
}

func (p *Pet) HappinessDecay(amount int32) {
	result := max(p.Happiness-amount, MinStat)
	atomic.StoreInt32(&p.Happiness, result)
}

func (p *Pet) Clean() {
	p.ShitView = 0
}

func (p *Pet) Die() {
	if p.Satisfaction <= 0 || p.Happiness <= 0 || p.ShitView > MaxStat {
		p.Dead = true
	}
}
