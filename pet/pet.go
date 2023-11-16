package pet

import (
	"math"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
	"virtualpet/location"
)

const (
	MaxStat int32 = 255
	MinStat int32 = 0

	WasteCoefficient float64 = 0.8
	MaxPoopLength    int32   = 50
)

type IPetActions interface {
	Feed(amount int32)
	Clean()
	Play()
	Travel()
}

type IPetDecay interface {
	HappinessDecay(amount int32)
	Starve(amount int32)
	Poop()
	Die()
}

type IPetChecks interface {
	GetPet() Pet
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
	PoopBuffer   int32
	PoopView     int32
	Location     location.Location
	Birthdate    time.Time
	LastFed      time.Time
	Dead         bool
}

var names = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Kathy", "Leo", "Molly", "Nathan", "Olivia", "Paul", "Quincy", "Rachel", "Sam", "Tom", "Ursula", "Victor", "Wendy", "Xander", "Yvonne", "Zane"}
var types = []string{"Dog", "Cat", "Bird", "Fish", "Bunny"}

func NewPet() IPet {
	p := Pet{
		Name:         names[rand.Intn(len(names))],
		Type:         types[rand.Intn(len(types))],
		Satisfaction: rand.Int31n(MaxStat),
		Happiness:    rand.Int31n(MaxStat),
		Location:     location.Home,
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

func (p *Pet) Travel() {
	var result string
	for i := 0; i < 10; i++ {
		result = location.locations[rand.Intn(len(location.locations))]
		if result != p.Location {
			break
		}
	}

	p.Location = result
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

func (p *Pet) GetPet() Pet {
	return *p
}

func (p *Pet) Poop() {
	if p.PoopBuffer == 0 {
		return
	}

	amount := min(p.PoopBuffer, rand.Int31n(MaxPoopLength))
	p.PoopView += amount
	p.PoopBuffer -= amount
}

func (p *Pet) Starve(amount int32) {
	result := max(p.Satisfaction-amount, MinStat)
	delta := p.Satisfaction - result
	waste := math.Abs(float64(delta) * WasteCoefficient)
	p.PoopBuffer += int32(waste)

	atomic.StoreInt32(&p.Satisfaction, result)
}

func (p *Pet) HappinessDecay(amount int32) {
	result := max(p.Happiness-amount, MinStat)
	atomic.StoreInt32(&p.Happiness, result)
}

func (p *Pet) Clean() {
	p.PoopView = 0
}

func (p *Pet) Die() {
	if p.Satisfaction <= 0 || p.Happiness <= 0 || p.PoopView > MaxStat {
		p.Dead = true
	}
}
