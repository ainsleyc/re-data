package redata 

import (
	"github.com/bitly/go-simplejson"
)

func ParseResults (data []byte) ([]byte, error) {
  respJson, _ := simplejson.NewJson(data)
  results, err := respJson.Get("map").Get("properties").EncodePretty() 

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

  return results, err 
}
