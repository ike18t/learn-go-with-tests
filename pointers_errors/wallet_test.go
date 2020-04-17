package wallet

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
  assert := assert.New(t)

  t.Run("Deposit", func(t *testing.T) {
    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))

    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("Withdraw", func(t *testing.T) {
    wallet := Wallet{Bitcoin(20)}

    err := wallet.Withdraw(Bitcoin(10))

    assert.NoError(err)
    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(20)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(100))

    assertBalance(t, wallet, startingBalance)
    assert.Equal(err, ErrInsufficientFunds)
  })
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
  t.Helper()

  got := wallet.Balance()

  if got != want {
    t.Errorf("got %s want %s", got, want)
  }
}
