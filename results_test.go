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
    { "$7.9K", 7900 },
    { "$327.89K", 327890 },
    { "$127.1119M", 127111000 },
  }

  for _, test := range tests {
    result := redata.ParsePriceString(test.input)
    if result != test.expected {
      t.Error(result, "!=", test.expected)
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
    result := redata.NormalizeDecimalString(test.input)
    if result != test.expected {
      t.Error(result, "!=", test.expected)
    }
  }
}

func TestNormalizeCoordinate(t *testing.T) {
  tests := []struct {
    input int64 
    expected float64 
  }{
    { 37834748, 37.834748 },
    { -127834748, -127.834748 },
  }

  for _, test := range tests {
    result := redata.NormalizeCoordinate(test.input)
    if result != test.expected {
      t.Error(result, "!=", test.expected)
    }
  }
}

