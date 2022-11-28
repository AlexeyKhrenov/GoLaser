package models

type Game struct {
	Round        Round
	GameScore    Score
	OverallScore Score
	Field        Field
}

func NewGame() Game {
	round := newRound()
	gameScore := Score{A: 0, B: 0}
	game := Game{Round: round, GameScore: gameScore, Field: fieldPrototype, OverallScore: overallScore}
	return game
}

func (g *Game) Shot(isSuccess bool) {
	shot := newShot(isSuccess)

	r := &g.Round
	r.shot(shot)
	if r.IsOver {
		g.OverallScore.add(r.IsSuccess)
		g.GameScore.add(r.IsSuccess)
	}
}

func (g *Game) NewRound() {
	round := newRound()
	round.start()
	g.Round = round
}
