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
    { "$359K", 359000 },
    { "$5M", 5000000 },
    // { "$327.89K", 327890 },
  }

  for _, test := range tests {
    price := redata.ParsePriceString(test.input)
    if price != test.expected {
      t.Error(price, "!=", test.expected)
    }
  }
}

func TestNormalizeDecimalString(t *testing.T) {
  tests := []struct {
    input string
    expected string 
  }{
    { "", "000" },
    { "3", "300" },
    { "56", "560" },
    { "678", "678" },
    { "2345", "234" },
  }

  for _, test := range tests {
    price := redata.NormalizeDecimalString(test.input)
    if price != test.expected {
      t.Error(price, "!=", test.expected)
    }
  }
}

