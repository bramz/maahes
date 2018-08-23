package commands

type Theo struct {
	output string
}

func RandomQuote() *Theo {
	return &Theo{
		output: string
	}
}

func (t *Theo) Register(name string) {
	t.output = name
	return
}