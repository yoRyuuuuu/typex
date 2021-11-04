package client

// Todo
type Game struct {
	problem string
}

func (g *Game) checkAnswer(input string) bool {
	return input == g.problem
}
