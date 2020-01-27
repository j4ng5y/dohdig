package main

import (
	"fmt"
	"log"

	"github.com/j4ng5y/dohdig/pkg/cloudflare"
	"github.com/j4ng5y/dohdig/pkg/common"
	"github.com/j4ng5y/dohdig/pkg/google"
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
		}
		providerFlag         string
		showOptionsFlag      bool
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
			Version: "0.2.0",
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

	gdigCmd.AddCommand(listCmd)
	gdigCmd.Flags().StringVarP(&providerFlag, "provider", "i", "google", "The provider to use")
	gdigCmd.Flags().StringVarP(&typeFlag, "record-type", "t", "A", "The DNS record type to query")
	gdigCmd.Flags().StringVarP(&ctFlag, "content-type", "c", "application/x-javascript", "The desired content type to return")
	gdigCmd.Flags().StringVarP(&eDNSClientSubnetFlag, "edns-client-subnet", "e", "0.0.0.0/0", "Set source IP address for DNS resolution")
	gdigCmd.Flags().StringVarP(&randomPaddingFlag, "random-padding", "p", "", "Pad request with random data")
	gdigCmd.Flags().BoolVarP(&cdFlag, "disable-dnssec-checking", "n", false, "Disable DNS validation")
	gdigCmd.Flags().BoolVarP(&doFlag, "show-dnssec", "d", true, "Show DNSSEC information in response")
	gdigCmd.Flags().BoolVarP(&showOptionsFlag, "show-options", "o", false, "Show configured options in the output")

	if err := gdigCmd.Execute(); err != nil {
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
