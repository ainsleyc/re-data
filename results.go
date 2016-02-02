package redata 

import (
  "strconv"
  "strings"
  "crypto/sha1"
  "encoding/json"
  "encoding/base32"
  "log"
  "regexp"

	"github.com/bitly/go-simplejson"
)

type PropertyResult struct {
  Id string
  PriceStr string
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
    p, _ := parseProperty(property.([]interface{}))
    results = append(results, p) 
  }

  resultsStr, _ := json.Marshal(results) 
  return resultsStr, nil 
}

func parseProperty (property []interface{}) (PropertyResult, error) {

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

  coordX, _ := property[0].(json.Number).Int64()
  coordY, _ := property[1].(json.Number).Int64()
  coordZ, _ := property[2].(json.Number).Int64()
  priceStr, _ := property[3].(string)
  coords := []string{
    strconv.FormatInt(coordX, 10),
    strconv.FormatInt(coordY, 10),
    strconv.FormatInt(coordZ, 10),
  }
  h := sha1.New()
  h.Write([]byte(strings.Join(coords, "")))
  id := strings.ToLower(base32.HexEncoding.EncodeToString((h.Sum(nil))))

  log.Println(id)

  return PropertyResult{
    id,
    priceStr,
    coordX,
    coordY,
    coordZ,
    4,
  }, nil
}

func ParsePriceString (price string) int {
  re := regexp.MustCompile("\\$(\\d+)([KM])")
  matches := re.FindAllStringSubmatch(price, -1)
  log.Println(matches)
  value, _ := strconv.Atoi(matches[0][1])
  if matches[0][2] == "K" {
    value = value * 1000
  }
  if matches[0][2] == "M" {
    value = value * 1000000
  }
  log.Println(value)
  return value 
}
