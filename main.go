package main

import (
	"fmt"
	"log"

	"github.com/j4ng5y/dohdig/pkg/blahdns"
	"github.com/j4ng5y/dohdig/pkg/cloudflare"
	"github.com/j4ng5y/dohdig/pkg/common"
	"github.com/j4ng5y/dohdig/pkg/google"
	"github.com/j4ng5y/dohdig/pkg/nextdns"
	"github.com/j4ng5y/dohdig/pkg/nixnet"
	"github.com/j4ng5y/dohdig/pkg/securedns"
	"github.com/j4ng5y/dohdig/pkg/snopyta"
	"github.com/spf13/cobra"
)

func execute() {
	var (
		validProviders = []string{
			"google",
			"cloudflare",
			"blahdns-fi",
			"blahdns-jp",
			"blahdns-de",
			"nextdns",
			"nixnet-uncensored",
			"nixnet-adblock",
			"nixnet-lasvegas",
			"nixnet-newyork",
			"nixnet-luxembourg",
			"securedns",
			"snopyta",
		}
		providerFlag         string
		showOptionsFlag      bool
		typeFlag             string
		cdFlag               bool
		ctFlag               string
		doFlag               bool
		eDNSClientSubnetFlag string
		randomPaddingFlag    string
		nextDNSID            string
		dohdigCmd            = &cobra.Command{
			Use:     "dohdig",
			Short:   "A small, dig-like command that only runs against the dns.google.com API",
			Example: "dohdig www.google.com",
			Version: "0.2.3",
			Args:    cobra.ExactArgs(1),
			Run: func(ccmd *cobra.Command, args []string) {
				var err error
				resp := new(common.QueryResponse)

				fmt.Printf("Querying: %s\n", args[0])
				if showOptionsFlag {
					fmt.Printf(
						optsStr,
						typeFlag,
						ctFlag,
						eDNSClientSubnetFlag,
						randomPaddingFlag,
						cdFlag,
						doFlag)
				}

				switch providerFlag {
				case "google":
					req := google.QueryRequest{
						Resource:                args[0],
						ResourceType:            typeFlag,
						ContentType:             ctFlag,
						EDNSClientSubnet:        eDNSClientSubnetFlag,
						RandomPadding:           randomPaddingFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "cloudflare":
					req := cloudflare.QueryRequest{
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "blahdns-fi":
					req := blahdns.QueryRequest{
						Country:                 "fi",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "blahdns-jp":
					req := blahdns.QueryRequest{
						Country:                 "jp",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "blahdns-de":
					req := blahdns.QueryRequest{
						Country:                 "de",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "nextdns":
					if nextDNSID == "" {
						log.Fatal("the --nextdns-id flag must be set to use NextDNS")
					}
					req := nextdns.QueryRequest{
						ID:                      nextDNSID,
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "nixnet-uncensored":
					req := nixnet.QueryRequest{
						ServerType:              "uncensored",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "nixnet-adblock":
					req := nixnet.QueryRequest{
						ServerType:              "adblock",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "nixnet-lasvegas":
					req := nixnet.QueryRequest{
						ServerType:              "lasvegas",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "nixnet-newyork":
					req := nixnet.QueryRequest{
						ServerType:              "newyork",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "nixnet-luxembourg":
					req := nixnet.QueryRequest{
						ServerType:              "luxembourg",
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "securedns":
					req := securedns.QueryRequest{
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				case "snopyta":
					req := snopyta.QueryRequest{
						Resource:                args[0],
						ResourceType:            typeFlag,
						DisableDNSSECValidation: cdFlag,
						ShowDNSSEC:              doFlag,
					}

					resp, err = req.Do()
					if err != nil {
						log.Fatal(err)
					}

					resp.Print()
				default:
					log.Fatalf("%s is an unsuppored provider", providerFlag)
				}
			},
		}

		listCmd = &cobra.Command{
			Use:   "list-providers",
			Short: "list available providers",
			Run: func(ccmd *cobra.Command, args []string) {
				fmt.Println("Valid Providers:")
				for _, v := range validProviders {
					fmt.Printf("  %s\n", v)
				}
			},
		}
	)

	dohdigCmd.AddCommand(listCmd)
	dohdigCmd.Flags().StringVarP(&providerFlag, "provider", "i", "google", "The provider to use")
	dohdigCmd.Flags().StringVarP(&typeFlag, "record-type", "t", "A", "The DNS record type to query")
	dohdigCmd.Flags().StringVarP(&ctFlag, "content-type", "c", "application/x-javascript", "The desired content type to return")
	dohdigCmd.Flags().StringVarP(&eDNSClientSubnetFlag, "edns-client-subnet", "e", "0.0.0.0/0", "Set source IP address for DNS resolution")
	dohdigCmd.Flags().StringVarP(&randomPaddingFlag, "random-padding", "p", "", "Pad request with random data")
	dohdigCmd.Flags().BoolVarP(&cdFlag, "disable-dnssec-checking", "n", false, "Disable DNS validation")
	dohdigCmd.Flags().BoolVarP(&doFlag, "show-dnssec", "d", true, "Show DNSSEC information in response")
	dohdigCmd.Flags().BoolVarP(&showOptionsFlag, "show-options", "o", false, "Show configured options in the output")

	if err := dohdigCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execute()
}

const optsStr string = `Options:
Record Type:        %s
Content Type:       %s
eDNS Client Subnet: %s
Random Pad:         %s
Disable DNSSEC:     %v
Show DNSSEC:        %v
`
