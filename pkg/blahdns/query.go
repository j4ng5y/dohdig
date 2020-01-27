package blahdns

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/j4ng5y/dohdig/pkg/common"
)

// QueryRequest is the request needed to query dns.google.com
type QueryRequest struct {
	Resource                string
	ResourceType            string
	DisableDNSSECValidation bool
	ShowDNSSEC              bool
}

// Do runs the query
//
// Arguments:
//     None
//
// Returns:
//     (*pkg.common.QueryResponse): A pointer to the query response, or nil if an error occurred
//     (error):                     An error if one exists, nil otherwise
func (q QueryRequest) Do(country string) (*common.QueryResponse, error) {
	var U string
	switch country {
	case "fi":
		U = fmt.Sprintf(
			"https://doh-fi.blahdns.com/dns-query?name=%s&type=%s&cd=%v&do=%v",
			q.Resource,
			q.ResourceType,
			q.DisableDNSSECValidation,
			q.ShowDNSSEC)
	case "jp":
		U = fmt.Sprintf(
			"https://doh-jp.blahdns.com/dns-query?name=%s&type=%s&cd=%v&do=%v",
			q.Resource,
			q.ResourceType,
			q.DisableDNSSECValidation,
			q.ShowDNSSEC)
	case "de":
		U = fmt.Sprintf(
			"https://doh-de.blahdns.com/dns-query?name=%s&type=%s&cd=%v&do=%v",
			q.Resource,
			q.ResourceType,
			q.DisableDNSSECValidation,
			q.ShowDNSSEC)
	default:
		return nil, fmt.Errorf("unsupported country")
	}

	u, err := url.Parse(U)
	if err != nil {
		return nil, fmt.Errorf("error parsing the provided url: %s, err: %w", q.Resource, err)
	}

	c := http.DefaultClient
	r, err := c.Do(&http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: http.Header{
			"accept": []string{
				"application/dns-json",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error sending the HTTP request, err: %w", err)
	}

	resp := new(common.QueryResponse)
	if err := resp.Unmarshal(r.Body); err != nil {
		return nil, fmt.Errorf("error unmarshalling the response, err: %w", err)
	}

	resp.DetermineStatusMessage()
	return resp, nil
}
