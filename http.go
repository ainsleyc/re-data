package redata 

import (
	"io/ioutil"
	"net/http"
	"net/url"

	// "github.com/asaskevich/govalidator"
)

func getResults(params url.Values) ([]byte, error) {
  baseUrl := "http://www.zillow.com/search/GetResults.htm"
  fullUrl := baseUrl + "?" + params.Encode()

  resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
  defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, ErrHttpBody(fullUrl)
	// }

	// if resp.StatusCode != 200 {
	// 	return nil, ErrHttpResponse(fullUrl, resp.StatusCode, body)
	// }

	return body, err 
}
