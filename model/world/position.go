package world

import "math"

// A three-dimensional, 64-bit integer position
type Position struct {
	X int64 // The x component of the position
	Y int64 // The y component of the position
	Z int64 // The z component of the position
}

// Calculates the distance between two positions
func CalculateDistance(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64((p2.X-p1.X)), 2) + math.Pow(float64((p2.Y-p1.Y)), 2) + math.Pow(float64((p2.Z-p1.Z)), 2))
}
