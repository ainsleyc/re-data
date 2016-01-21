package redata 

import (
	"github.com/bitly/go-simplejson"
)

func ParseResults (data []byte) ([]byte, error) {
  respJson, _ := simplejson.NewJson(data)
  results, err := respJson.EncodePretty() 
  return results, err 
}
