package mixpanel

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	ApiSecret  string
}

func (c *Client) Get(endpoint string, params url.Values) (*http.Response, error) {
	c.addSignature(params)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", endpoint, params.Encode()), nil)
	if err != nil {
		return nil, err
	}
	return c.HttpClient.Do(req)
}

func (c *Client) addSignature(params url.Values) {
	// add required params
	expire := time.Now().Add(time.Minute * 10).Unix()
	params.Add("expire", fmt.Sprintf("%v", expire))
	params.Add("api_key", c.ApiKey)
	var kvList []string
	// create key/value array
	for k, vs := range params {
		for _, v := range vs {
			kvList = append(kvList, fmt.Sprintf("%s=%s", k, v))
		}
	}
	// sort it
	sort.StringSlice(kvList).Sort()

	// join the sorted array and append the secret
	basestring := strings.Join(kvList, "") + c.ApiSecret
	//fmt.Printf("BS: %v\n", basestring)
	// calculate the md5 sum and base64 encode it
	hsh := md5.New()
	hsh.Write([]byte(basestring))
	data := hsh.Sum(nil)
	sig := hex.EncodeToString(data)
	//fmt.Printf("Sig: %v\n", sig)
	params.Add("sig", sig)
}
