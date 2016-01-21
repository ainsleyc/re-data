package redata 

import (
	"io/ioutil"
	"net/http"
	"net/url"

  "golang.org/x/net/context"
  "google.golang.org/appengine"
  "google.golang.org/appengine/urlfetch"
)

type Client struct {
  context context.Context
  client *http.Client
}

func NewClient(r *http.Request) Client {
  ctx := appengine.NewContext(r)
  client := urlfetch.Client(ctx)
  return Client{ctx, client}
}

func (c Client) GetResults(params url.Values) ([]byte, error) {
  baseUrl := "http://www.zillow.com/search/GetResults.htm"
  fullUrl := baseUrl + "?" + params.Encode()

  resp, err := c.client.Get(fullUrl)
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
