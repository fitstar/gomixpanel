Mixpanel Export Client
======================

Simple client that performs the mixpanel signing for the export API documented:

https://mixpanel.com/docs/api-documentation/data-export-api

Auth documented here:

https://mixpanel.com/docs/api-documentation/data-export-api#auth-implementation

example:

    import "github.com/fitstar/mixpanel"

    c := &mixpanel.Client{
		HttpClient: http.DefaultClient,
		ApiKey:     "my_key",
		ApiSecret:  "my_secret",
	}
    params := make(url.Values)
    params.Add("from_date", "2013-11-01")
    params.Add("to_date", "2013-11-02")

	resp, err := c.Get("https://data.mixpanel.com/api/2.0/export/", params)

Mixpanel throttles (resp.StatusCode == 429) aggressively so don't bother with goroutines.
