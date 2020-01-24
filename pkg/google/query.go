package google

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/j4ng5y/gdig/pkg/common"
)

// QueryRequest is the request needed to query dns.google.com
type QueryRequest struct {
	Resource                string
	ResourceType            string
	ContentType             string
	EDNSClientSubnet        string
	RandomPadding           string
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
func (q QueryRequest) Do() (*common.QueryResponse, error) {
	U := fmt.Sprintf(
		"https://dns.google.com/resolve?name=%s&type=%s&ct=%s&edns_client_subnet=%s&cd=%v&do=%v",
		q.Resource,
		q.ResourceType,
		q.ContentType,
		q.EDNSClientSubnet,
		q.DisableDNSSECValidation,
		q.ShowDNSSEC)
	u, err := url.Parse(U)
	if err != nil {
		return nil, fmt.Errorf("error parsing the provided url: %s, err: %w", q.Resource, err)
	}

	c := http.DefaultClient
	r, err := c.Do(&http.Request{
		Method: http.MethodGet,
		URL:    u,
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
