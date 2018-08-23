package commands

type Test struct {
	output string
}

func TestMsg() *Test {
	return &Test{
		output: string
	}
}

func (t *Test) Register(name string) {
	t.output = name
	return
}
