package game

func (g *Game) withinBounds(p position) bool {
	x, y := g.state.GetDimensions()
	if p.x < 0 || p.x > x {
		return false
	}
	if p.y < 0 || p.y > y {
		return false
	}
	return true
}
