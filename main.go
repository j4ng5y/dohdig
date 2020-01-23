package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

type queryRequest struct {
	resource         string
	resourceType     string
	contentType      string
	eDNSClientSubnet string
	randomPadding    string
	disableDNSSEC    bool
	showDNSSEC       bool
}

type queryResponse struct {
	Status        int `json:"Status"`
	statusMessage string
	TC            bool `json:"TC"`
	RD            bool `json:"RD"`
	RA            bool `json:"RA"`
	AD            bool `json:"AD"`
	CD            bool `json:"CD"`
	Question      []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"Question"`
	Answer []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		TTL  int    `json:"TTL"`
		Data string `json:"data"`
	} `json:"Answer"`
	Additional       []interface{} `json:"Additional"`
	EDNSClientSubnet string        `json:"edns_client_subnet"`
}

func (q *queryResponse) unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, q)
}

func (q queryRequest) do() (*queryResponse, error) {
	u, err := url.Parse(
		fmt.Sprintf(
			"https://dns.google.com/resolve?name=%s&type=%s&ct=%s&edns_client_subnet=%s&cd=%v&do=%v",
			q.resource,
			q.resourceType,
			q.contentType,
			q.eDNSClientSubnet,
			q.disableDNSSEC,
			q.showDNSSEC))
	if err != nil {
		return nil, fmt.Errorf("error parsing the provided url: %s, err: %w", q.resource, err)
	}

	c := http.DefaultClient
	r, err := c.Do(&http.Request{
		Method: http.MethodGet,
		URL:    u,
	})
	if err != nil {
		return nil, fmt.Errorf("error sending the HTTP request, err: %w", err)
	}

	resp := new(queryResponse)
	if err := resp.unmarshal(r.Body); err != nil {
		return nil, fmt.Errorf("error unmarshalling the response, err: %w", err)
	}

	switch resp.Status {
	case 0:
		resp.statusMessage = "NOERROR: No Errors Reported"
	case 1:
		resp.statusMessage = "FORMERR: The DNS Query Is Malformed"
	case 2:
		resp.statusMessage = "SERVFAIL: The DNS Server Failed To Process This Request"
	case 3:
		resp.statusMessage = "NXDOMAIN: The Requested Domain Name Does Not Exist"
	case 4:
		resp.statusMessage = "NOTIMP: This is not implimented"
	case 5:
		resp.statusMessage = "REFUSED: The DNS Server Refused To Answer This Query"
	case 6:
		resp.statusMessage = "YXDOMAIN: The Requested Domain Exists, But It Should Not"
	case 7:
		resp.statusMessage = "XRRSET: The Requested Resource Record Set Exists, But It Should Not"
	case 8:
		resp.statusMessage = "NOTAUTH: The DNS Server Is Not Authorized To Respond To Queries For This DNS Zone"
	case 9:
		resp.statusMessage = "NOTZONE: The Requested DNS Name Does Not Exist In The DNS Servers Authorized Zones"
	}

	return resp, nil
}

func execute() {
	var (
		typeFlag             string
		cdFlag               bool
		ctFlag               string
		doFlag               bool
		eDNSClientSubnetFlag string
		randomPaddingFlag    string
		gdigCmd              = &cobra.Command{
			Use:     "gdig",
			Short:   "A small, dig-like command that only runs against the dns.google.com API",
			Example: "gdig www.google.com",
			Version: "0.1.0",
			Args:    cobra.ExactArgs(1),
			Run: func(ccmd *cobra.Command, args []string) {
				str := `Querying:
  %s
Options:
  Record Type:        %s
  Content Type:       %s
  eDNS Client Subnet: %s
  Random Pad:         %s
  Disable DNSSEC:     %v
  Show DNSSEC:        %v
Answer:
  Status:             %s
  Truncated:          %v
  RD:                 %v
  RA:                 %v
  DNSSEC Validated:   %v
  DNSSEC Disabled:    %v
  eDNS Client Subnet: %s
  Data:`
				req := queryRequest{
					resource:         args[0],
					resourceType:     typeFlag,
					contentType:      ctFlag,
					eDNSClientSubnet: eDNSClientSubnetFlag,
					randomPadding:    randomPaddingFlag,
					disableDNSSEC:    cdFlag,
					showDNSSEC:       doFlag,
				}
				resp, err := req.do()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(
					fmt.Sprintf(
						str,
						args[0],
						typeFlag,
						ctFlag,
						eDNSClientSubnetFlag,
						randomPaddingFlag,
						cdFlag,
						doFlag,
						resp.statusMessage,
						resp.TC,
						resp.RD,
						resp.RA,
						resp.AD,
						resp.CD,
						resp.EDNSClientSubnet))
				for _, s := range resp.Answer {
					fmt.Println(fmt.Sprintf("    %s", s.Data))
				}
				for _, s := range resp.Additional {
					fmt.Println(fmt.Sprintf("Additional: %v", s))
				}
			},
		}
	)

	gdigCmd.Flags().StringVarP(&typeFlag, "record-type", "t", "A", "The DNS record type to query")
	gdigCmd.Flags().StringVarP(&ctFlag, "content-type", "c", "application/x-javascript", "The desired content type to return")
	gdigCmd.Flags().StringVarP(&eDNSClientSubnetFlag, "edns-client-subnet", "e", "0.0.0.0/0", "Set source IP address for DNS resolution")
	gdigCmd.Flags().StringVarP(&randomPaddingFlag, "random-padding", "p", "", "Pad request with random data")
	gdigCmd.Flags().BoolVarP(&cdFlag, "disable-dnssec-checking", "n", false, "Disable DNS validation")
	gdigCmd.Flags().BoolVarP(&doFlag, "show-dnssec", "d", true, "Show DNSSEC information in response")

	if err := gdigCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execute()
}
