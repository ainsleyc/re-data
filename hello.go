package redata 

import (
    "fmt"
    "net/http"
	  "net/url"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    client := NewClient(r)

    params := url.Values{}
    params.Set("spt", "homes")
    params.Set("status", "110001")
    params.Set("lt", "111101")
    params.Set("ht", "111111")
    params.Set("pr", ",")
    params.Set("mp", ",")
    params.Set("bd", "1%2C")
    params.Set("ba", "0%2C")
    params.Set("sf", ",")
    params.Set("lot", ",")
    params.Set("yr", ",")
    params.Set("pho", "0")
    params.Set("pets", "0")
    params.Set("parking", "0")
    params.Set("laundry", "0")
    params.Set("pnd", "0")
    params.Set("red", "0")
    params.Set("zso", "0")
    params.Set("days", "any")
    params.Set("ds", "all")
    params.Set("pmf", "1")
    params.Set("pf", "1")
    params.Set("zoom", "10")
    params.Set("rect", "-122414131,37301095,-121521492,37439156")
    params.Set("p", "1")
    params.Set("sort", "days")
    params.Set("search", "maplist")
    params.Set("disp", "1")
    params.Set("rid", "13713")
    params.Set("rt", "6")
    params.Set("listright", "true")
    params.Set("isMapSearch", "1")
    params.Set("zoom", "10")

    body, err := client.GetResults(params)
    data, _ := ParseResults(body)
    fmt.Println(w, err)
    fmt.Fprint(w, string(data))
}
