package commands

func TheoCmd {
	theo, err := ioutil.ReadFile("data/theo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(theo), "\n")

	if content == "!theo" {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		i := r.Intn(len(lines) - 1)
		line := lines[i]
		return line
}