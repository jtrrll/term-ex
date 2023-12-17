package position

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSurrounding(t *testing.T) {
	for _, inputs := range [][4]int{
		{0, 0, 0, 0},
		{4243, 1513, 1, 4},
		{024, 25801, 2, 12},
		{10513, 524, 3, 28},
		{3, 994, 4, 48}} {
		position := Position{int64(inputs[0]), int64(inputs[1])}
		actual := position.GetSurrounding(uint64(inputs[2]))
		assert.NotContains(t, actual, position)
		assert.Len(t, actual, inputs[3])
	}
}

func FuzzGetSurrounding(f *testing.F) {
	f.Add(int64(0), int64(0), uint64(0))
	f.Add(int64(0), int64(0), uint64(1))
	f.Add(int64(45), int64(-13), uint64(3))

	f.Fuzz(func(t *testing.T, x int64, y int64, radius uint64) {
		center := Position{x, y}
		actual := center.GetSurrounding(radius)
		assert.NotContains(t, actual, center)
		assert.LessOrEqual(t, len(actual), int(math.Pi*math.Pow(float64(radius)+math.Sqrt2, 2)))
		for _, v := range actual {
			assert.LessOrEqual(t, CalculateDistance(center, v), float64(radius))
		}
	})
}

func TestCalculateDistance(t *testing.T) {
	for _, inputs := range [][5]int64{
		{0, 0, 0, 0, 0},
		{5, 0, 11, 0, 6},
		{20, -5, 23, -1, 5}} {
		p1 := Position{inputs[0], inputs[1]}
		p2 := Position{inputs[2], inputs[3]}
		assert.Equal(t, float64(inputs[4]), CalculateDistance(p1, p2))
		assert.Equal(t, float64(inputs[4]), CalculateDistance(p2, p1))
	}
}

func FuzzCalculateDistance(f *testing.F) {
	f.Add(int64(0), int64(0), int64(0), int64(0))
	f.Add(int64(42), int64(241), int64(-153), int64(-135))
	f.Add(int64(15153), int64(-5113), int64(23443), int64(-1002))

	f.Fuzz(func(t *testing.T, x1 int64, y1 int64, x2 int64, y2 int64) {
		p1 := Position{x1, y1}
		p2 := Position{x2, y2}

		xDist := x2 - x1
		yDist := y2 - y1
		expected := math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
		assert.InDelta(t, expected, CalculateDistance(p1, p2), .001)
		assert.InDelta(t, expected, CalculateDistance(p2, p1), .001)
		assert.InDelta(t, math.Pow(expected, 2), math.Pow(float64(xDist), 2)+math.Pow(float64(yDist), 2), .001)
	})
}
