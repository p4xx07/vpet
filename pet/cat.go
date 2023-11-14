package pet

type Cat struct {
	Pet
}

func (p *Cat) GetFrames() []string {
	return []string{
		`/\_/\  
( o.o ) 
 > ^ <  
	`,
		` /\_/\  
( -.- ) 
 > ^ <
`,
		` /\_/\  
( o.o ) 
 > ~ <
`,
		`/ /\_/\  
( -.- ) 
 > ~ <
`,
	}
}
