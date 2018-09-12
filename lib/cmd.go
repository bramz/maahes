package lib

type Cmd interface {
	Handle([]string) string
}
