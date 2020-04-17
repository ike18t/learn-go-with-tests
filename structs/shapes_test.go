package structs

import (
  "reflect"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
  assert := assert.New(t)

  got := Perimeter(Rectangle{10.0, 10.0})
  want := 40.0

  assert.Equal(want, got)
}

func TestArea(t *testing.T) {
  assert := assert.New(t)

  areaTests := []struct {
    shape Shape
    want float64
  }{
    {shape: Rectangle{Width: 12, Height: 6.0}, want: 72.0},
    {shape: Circle{Radius: 10}, want: 314.1592653589793},
    {shape: Triangle{Base: 12, Height: 6}, want: 36.0},
  }

  for _, tt := range areaTests {
    t.Run(reflect.TypeOf(tt.shape).Name(), func(t *testing.T) {
      assert.Equal(tt.want, tt.shape.Area())
    })
  }
}
