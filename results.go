package redata 

import (
  "strconv"
  "strings"
  "crypto/sha1"
  "encoding/json"
  "encoding/base32"
  "unicode/utf8"
  "log"
  "regexp"

	"github.com/bitly/go-simplejson"
)

type PropertyResult struct {
  Id string
  Price int 
  Beds int64
  Baths float64 
  SqFt int64
  SqFtStr string
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

  subArray, _ := property[7].([]interface{})
  log.Println(subArray)
  beds, _ := subArray[1].(json.Number).Int64()
  baths, _ := subArray[2].(json.Number).Float64()
  sqFtStr, _ := subArray[6].(string)

  return PropertyResult{
    id,
    ParsePriceString(priceStr),
    beds,
    baths,
    4,
    sqFtStr,
  }, nil
}

func ParsePriceString (price string) int {
  re := regexp.MustCompile("\\$(\\d+)\\.?(\\d+)?([KM])")
  matches := re.FindAllStringSubmatch(price, -1)
  value, _ := strconv.Atoi(matches[0][1])
  decimalStr := NormalizeDecimalString(matches[0][2])
  decimal, _ := strconv.Atoi(decimalStr)
  if matches[0][3] == "K" {
    value = value * 1000 + decimal
  }
  if matches[0][3] == "M" {
    value = value * 1000000 + decimal * 1000
  }
  return value 
}

func NormalizeDecimalString (decimal string) string {
  len := utf8.RuneCountInString(decimal)
  switch {
  case len == 0:
    return "000"
  case len == 1:
    return decimal + "00"
  case len == 2:
    return decimal + "0"
  case len > 3:
    return decimal[0:3]
  }
  return decimal
}
