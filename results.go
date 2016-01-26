package redata 

import (
  "fmt"
  "encoding/json"

	"github.com/bitly/go-simplejson"
)

type PropertyResult struct {
  Id string
  Price int64 
  Beds int64
  Baths int64
  SqFt int64
}

func ParseResults (data []byte) ([]byte, error) {
  respJson, _ := simplejson.NewJson(data)
  properties := respJson.Get("map").Get("properties")

  results := []PropertyResult{} 
  for _, property := range properties.MustArray() {
    propertyArray := property.([]interface{})
    coordX, _ := propertyArray[0].(json.Number).Int64()
    coordY, _ := propertyArray[1].(json.Number).Int64()
    coordZ, _ := propertyArray[2].(json.Number).Int64()
    fmt.Println(propertyArray)
    results = append(results, PropertyResult{
      "blah",
      coordX,
      coordY,
      coordZ,
      4,
    })
  }

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

  // test, err := properties.EncodePretty()
  // return test, err 

  resultsStr, _ := json.Marshal(results) 
  return resultsStr, nil 
}
