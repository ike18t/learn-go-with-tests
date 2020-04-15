package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
  assert := assert.New(t)

  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Ike", "")
    want := "Hello, Ike"

    assert.Equal(want, got)
  })

  t.Run("empty string defaults to 'World'", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"

    assert.Equal(want, got)
  })

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Ike", "Spanish")
    want := "Hola, Ike"

    assert.Equal(want, got)
  })

  t.Run("in French", func(t *testing.T) {
    got := Hello("Ike", "French")
    want := "Bonjour, Ike"

    assert.Equal(want, got)
  })
}
