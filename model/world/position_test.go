package world

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzGetNeighbors(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(3), int64(-4))
	f.Add(int64(-100), int64(777))

	f.Fuzz(func(t *testing.T, x int64, y int64) {
		actual := Position{x, y}.GetNeighbors()
		assert.Contains(t, actual, Position{x, y - 1}) // north
		assert.Contains(t, actual, Position{x, y + 1}) // south
		assert.Contains(t, actual, Position{x - 1, y}) // east
		assert.Contains(t, actual, Position{x + 1, y}) // west
	})
}

func FuzzCalculateDistance(f *testing.F) {
	f.Add(int64(0), int64(0), int64(0), int64(0))
	f.Add(int64(5), int64(0), int64(11), int64(0))
	f.Add(int64(20), int64(-5), int64(23), int64(-1))

	f.Fuzz(func(t *testing.T, x1 int64, y1 int64, x2 int64, y2 int64) {
		p1 := Position{x1, y1}
		p2 := Position{x2, y2}

		xDist := x2 - x1
		yDist := y2 - y1
		actual := math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
		assert.Equal(t, actual, CalculateDistance(p1, p2))
		assert.Equal(t, actual, CalculateDistance(p2, p1))
	})
}
