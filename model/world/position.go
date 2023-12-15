package world

import "math"

// A two-dimensional, 64-bit integer position
type Position struct {
	X int64 // The x component of the position
	Y int64 // The y component of the position
}

// Returns a list of the position's neighbors
func (p Position) GetNeighbors() [4]Position {
	return [4]Position{
		{p.X, p.Y - 1}, // north
		{p.X, p.Y + 1}, // south
		{p.X - 1, p.Y}, // east
		{p.X + 1, p.Y}, // west
	}
}

// Calculates the distance between two positions
func CalculateDistance(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64((p2.X-p1.X)), 2) + math.Pow(float64((p2.Y-p1.Y)), 2))
}
