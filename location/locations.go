package location

var locations = []string{"home", "park", "supermarket", "office", "library", "school", "hospital", "restaurant", "gym"}

type Location struct {
	Name                 string
	HappinessCoefficient float32
	PoopCoefficient      float32
	PoopQuantity         int32
	ASCIIArt             string
}

var (
	Home = Location{
		Name:                 "home",
		HappinessCoefficient: 0.5,
		PoopCoefficient:      1,
		PoopQuantity:         0,
		ASCIIArt:             "ascii of a home",
	}
)
