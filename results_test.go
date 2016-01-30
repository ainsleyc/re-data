package redata_test

import (
  "testing"

  "github.com/ainsleyc/redata"
)

func TestParsePriceString_ShouldReturnCorrectResult(t *testing.T) {
  tests := []struct {
    input string
    expected int
  }{
    { "$5M", 5000000 },
  }

  for _, test := range tests {
    price := redata.ParsePriceString(test.input)
    if price != test.expected {
      t.Error(price, "!=", test.expected)
    }
  }
}

