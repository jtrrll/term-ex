package set

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	actual1 := Intersection(mapset.NewSet(4, 5, 2), mapset.NewSet(4, 2, 8), mapset.NewSet(1, 2, 3, 4))
	expected1 := mapset.NewSet(2, 4)
	assert.Equal(t, expected1, actual1)

	actual2 := Intersection(mapset.NewSet[int](), mapset.NewSet(5, 4, 3, 2, 1, 0))
	expected2 := mapset.NewSet[int]()
	assert.Equal(t, expected2, actual2)

	actual3 := Intersection(mapset.NewSet("", "abcd", "123"), mapset.NewSet("", "abcd", "123"), mapset.NewSet("", "abcd", "123"), mapset.NewSet("", "abcd", "123"))
	expected3 := mapset.NewSet("", "abcd", "123")
	assert.Equal(t, expected3, actual3)
}

func FuzzIntersection(f *testing.F) {
	f.Add(uint(0), uint(0), uint(0))
	f.Add(uint(125), uint(10), uint(4))
	f.Add(uint(77), uint(130), uint(46))

	f.Fuzz(func(t *testing.T, count1 uint, count2 uint, count3 uint) {
		s1 := mapset.NewSet[int]()
		for i := 0; i < int(count1); i++ {
			s1.Add(i)
		}
		assert.Equal(t, int(count1), s1.Cardinality())
		s2 := mapset.NewSet[int]()
		for i := 0; i < int(count2); i++ {
			s2.Add(i)
		}
		assert.Equal(t, int(count2), s2.Cardinality())
		s3 := mapset.NewSet[int]()
		for i := 0; i < int(count3); i++ {
			s3.Add(i)
		}
		assert.Equal(t, int(count3), s3.Cardinality())

		actual := Intersection(s1, s2, s3)
		minCount := min(count1, count2, count3)
		assert.Equal(t, int(minCount), actual.Cardinality())
		for i := 0; i < int(minCount); i++ {
			assert.True(t, actual.Contains(i))
		}
	})
}
