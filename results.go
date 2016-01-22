package redata 

import (
	"github.com/bitly/go-simplejson"
)

func ParseResults (data []byte) ([]byte, error) {
  respJson, _ := simplejson.NewJson(data)
  results, err := respJson.Get("map").Get("properties").EncodePretty() 
  return results, err 
}
