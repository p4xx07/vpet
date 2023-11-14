package pet

type Bird struct {
	Pet
}

func (p *Bird) GetFrames() []string {
	return []string{
		`  \ /
 ( o.o )
    >
   / \
	`,
		`  \ /
 ( o.o )
    <
   / \
`,
		`  \ /
 ( o.o )
    >
   \_/
`,
		`  \ /
 ( o.o )
    <
   \_/
`,
	}
}
