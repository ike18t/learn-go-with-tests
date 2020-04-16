package arrays

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
  assert := assert.New(t)

  t.Run("collection of any size", func(t *testing.T) {
    numbers := []int{1, 2, 3}

    got := Sum(numbers)
    want := 6

    assert.Equal(want, got)
  })
}

func TestSumAll(t *testing.T) {
  assert := assert.New(t)

  got := SumAll([]int{1,2}, []int{0,9})
  want := []int{3, 9}

  assert.Equal(want, got)
}

func TestSumAllTails(t *testing.T) {
  assert := assert.New(t)

  t.Run("make the sums of some slices", func(t *testing.T) {
    got := SumAllTails([]int{1,2}, []int{0,9})
    want := []int{2, 9}

    assert.Equal(want, got)
  })

  t.Run("safely sum empty slices", func(t *testing.T) {
    got := SumAllTails([]int{}, []int{3, 4, 5})
    want := []int{0, 9}

    assert.Equal(want, got)
  })
}

