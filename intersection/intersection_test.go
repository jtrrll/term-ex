package intersection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	actual1 := Intersection([]int{4, 5, 2}, []int{4, 2, 8}, []int{1, 2, 3, 4})
	assert.Len(t, actual1, 2)
	assert.Contains(t, actual1, 2)
	assert.Contains(t, actual1, 4)

	actual2 := Intersection([]int{}, []int{5, 4, 3, 2, 1, 0})
	assert.Len(t, actual2, 0)

	actual3 := Intersection([]string{"", "abcd", "123"}, []string{"", "abcd", "123"}, []string{"", "abcd", "123"}, []string{"", "abcd", "123"})
	assert.Len(t, actual3, 3)
	assert.Contains(t, actual3, "")
	assert.Contains(t, actual3, "abcd")
	assert.Contains(t, actual3, "123")
}
