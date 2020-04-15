package integers

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
  assert := assert.New(t)

  sum := Add(2, 2)
  expected := 4

  assert.Equal(expected, sum)
}

func ExampleAdd() {
  sum := Add(1, 5)
  fmt.Println(sum)
  // Output: 6
}
