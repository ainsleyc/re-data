package redata 

import (
	"github.com/bitly/go-simplejson"
)

type PropertyResult struct {
  Id string
  Price int 
  Beds int
  Baths int
  SqFt int
}

func ParseResults (data []byte) ([]byte, error) {
  respJson, _ := simplejson.NewJson(data)
  properties := respJson.Get("map").Get("properties")

  // results := []PropertyResult{} 
  // for _, property := range properties.MustArray() {
  //   results = append(results, PropertyResult{
  //     "blah",
  //     1,
  //     2,
  //     3,
  //     4,
  //   })
  // }

  // data structure
  // [
  //   Location X 
  //   Location Y 
  //   Location Z 
  //   Price ($ K/?)
  //   ?
  //   ?
  //   ?
  //   [
  //     Price ($ K/?)
  //     Beds
  //     Baths
  //     Square Feet 
  //     (seems to be always false) 
  //     Image URL
  //     Size string
  //     [
  //       Extra Information Type
  //         1 = Open House Time
  //         ...
  //     ]
  //   ]
  // ]

  test, err := properties.EncodePretty()
  return test, err 
}
