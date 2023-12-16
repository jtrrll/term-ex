package world

import "math"

// A two-dimensional, 64-bit integer position
type Position struct {
	X int64 // The x component of the position
	Y int64 // The y component of the position
}

// Returns a set of positions within a radius around this position
func (p Position) GetSurrounding(radius uint64) []Position {
	positions := []Position{}
	for x := p.X - int64(radius); x <= p.X+int64(radius); x++ {
		for y := p.Y - int64(radius); y <= p.Y+int64(radius); y++ {
			if x != p.X || y != p.Y {
				position := Position{x, y}
				if CalculateDistance(p, position) <= float64(radius) {
					positions = append(positions, position)
				}
			}
		}
	}
	return positions
}

// Calculates the distance between two positions
func CalculateDistance(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64((p2.X-p1.X)), 2) + math.Pow(float64((p2.Y-p1.Y)), 2))
}
