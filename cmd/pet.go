package main

type Pet struct {
	hungerLevel    int
	happinessLevel int
}

func (p *Pet) Feed() {
	p.hungerLevel -= 1
}

func (p *Pet) Play() {
	p.happinessLevel += 1
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
