package pet

type Dog struct {
	Pet
}

func (p *Dog) GetFrames() []string {
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
