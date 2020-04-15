package iteration

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
  assert := assert.New(t)

  t.Run("takes arg of the char to repeat and how many times", func(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    assert.Equal(repeated, expected)
  })
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 5)
  }
}
