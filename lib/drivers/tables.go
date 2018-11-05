package drivers

type Messages struct {
	Id            int
	SendTime      string
	Content       string
	User          string
	Discriminator string
	Channel       string
	Server        string
}

type Quotes struct {
	Id    int
	DTime string
	Quote string
}
