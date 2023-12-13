package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistance(t *testing.T) {
	p1 := Position{}
	p2 := Position{}
	assert.EqualValues(t, 0, CalculateDistance(p1, p2))
	assert.EqualValues(t, 0, CalculateDistance(p2, p1))

	p1 = Position{X: 5}
	p2 = Position{X: 11}
	assert.EqualValues(t, 6, CalculateDistance(p1, p2))
	assert.EqualValues(t, 6, CalculateDistance(p2, p1))

	p1 = Position{X: 20, Y: -5}
	p2 = Position{X: 23, Y: -1}
	assert.EqualValues(t, 5, CalculateDistance(p1, p2))
	assert.EqualValues(t, 5, CalculateDistance(p2, p1))

	p1 = Position{X: 6, Y: 4, Z: -3}
	p2 = Position{X: 2, Y: -8, Z: 3}
	assert.EqualValues(t, 14, CalculateDistance(p1, p2))
	assert.EqualValues(t, 14, CalculateDistance(p2, p1))
}
