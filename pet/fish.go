package pet

type Fish struct {
	Pet
}

func (p *Fish) GetFrames() []string {
	return []string{
		`><(((º>
	`,
		`<º)))><
`,
		`><(((º>
`,
		`<º)))><
`,
	}
}
