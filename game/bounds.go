package game

func (g *Game) withinBounds(p position) bool {
	x, y := g.state.GetDimensions()
	if p.X < 0 || p.X > x {
		return false
	}
	if p.Y < 0 || p.Y > y {
		return false
	}
	return true
}
