package pet

type Bunny struct {
	Pet
}

func (p *Bunny) GetFrames() []string {
	return []string{
		`(\__/)
(='.'=)
(")_(")
	`,
		`(\__/)
(='.'=)
(")_(")
`,
		`(\__/)
(='.'=)
(")_(")
`,
		`(\__/)
(='.'=)
(")_(")
`,
	}
}
