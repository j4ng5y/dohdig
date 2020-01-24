package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Do is a standard interface for running queries
type Do interface {
	Do() (*QueryResponse, error)
}

// QueryResponseQuestion - Question struct for the QueryResponse struct
type QueryResponseQuestion struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

// QueryResponseAnswer - Answer struct for the QueryResponse struct
type QueryResponseAnswer struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	TypeName    string `json:"-"`
	TypeMeaning string `json:"-"`
	TTL         int    `json:"TTL"`
	Data        string `json:"data"`
}

// DetermineTypeNameAndMeaning will generate the string Name and Meaning for the Provided Type
//
// Arguments:
//     None
//
// Returns:
//     None
func (q *QueryResponseAnswer) DetermineTypeNameAndMeaning() {
	switch q.Type {
	case 1:
		q.TypeName = "A"
		q.TypeMeaning = "A Host Address"
	case 2:
		q.TypeName = "NS"
		q.TypeMeaning = "An Authoritative Name Server"
	case 3:
		q.TypeName = "MD"
		q.TypeMeaning = "A Mail Destination"
	case 4:
		q.TypeName = "MF"
		q.TypeMeaning = "A Mail Forwarder"
	case 5:
		q.TypeName = "CNAME"
		q.TypeMeaning = "The Canonical Name For An Alias"
	case 6:
		q.TypeName = "SOA"
		q.TypeMeaning = "Marks The Start Of A Zone Of Authority"
	case 7:
		q.TypeName = "MB"
		q.TypeMeaning = "A Mailbox Domain Name"
	case 8:
		q.TypeName = "MG"
		q.TypeMeaning = "A Mail Group Member"
	case 9:
		q.TypeName = "MR"
		q.TypeMeaning = "A Mail Rename Domain Name"
	case 10:
		q.TypeName = "NULL"
		q.TypeMeaning = "A Null Resource Record"
	case 11:
		q.TypeName = "WKS"
		q.TypeMeaning = "A Well Known Service Description"
	case 12:
		q.TypeName = "PTR"
		q.TypeMeaning = "A Domain Name Pointer"
	case 13:
		q.TypeName = "HINFO"
		q.TypeMeaning = "Host Information"
	case 14:
		q.TypeName = "MINFO"
		q.TypeMeaning = "Mailbox Or Mail List Information"
	case 15:
		q.TypeName = "MX"
		q.TypeMeaning = "Mail Exchange"
	case 16:
		q.TypeName = "TXT"
		q.TypeMeaning = "Text Strings"
	case 17:
		q.TypeName = "RP"
		q.TypeMeaning = "For Responsible Person"
	case 18:
		q.TypeName = "AFSDB"
		q.TypeMeaning = "For AFS Data Base Location"
	case 19:
		q.TypeName = "X25"
		q.TypeMeaning = "For X.25 PSDN Address"
	case 20:
		q.TypeName = "ISDN"
		q.TypeMeaning = "For ISDN Address"
	case 21:
		q.TypeName = "RT"
		q.TypeMeaning = "For Route Through"
	case 22:
		q.TypeName = "NSAP"
		q.TypeMeaning = "For NSAP Address, NSAP Style A Record"
	case 23:
		q.TypeName = "NSAP-PTR"
		q.TypeMeaning = "For Domain Name Pointer, NSAP Style"
	case 24:
		q.TypeName = "SIG"
		q.TypeMeaning = "For Security Signature"
	case 25:
		q.TypeName = "KEY"
		q.TypeMeaning = "For Security Key"
	case 26:
		q.TypeName = "PX"
		q.TypeMeaning = "X.400 Mail Mapping Information"
	case 27:
		q.TypeName = "GPOS"
		q.TypeMeaning = "Geographical Position"
	case 28:
		q.TypeName = "AAAA"
		q.TypeMeaning = "IPV6 Address"
	case 29:
		q.TypeName = "LOC"
		q.TypeMeaning = "Location Information"
	case 30:
		q.TypeName = "NXT"
		q.TypeMeaning = "Next Domain"
	case 31:
		q.TypeName = "EID"
		q.TypeMeaning = "Endpoint Identifier"
	case 32:
		q.TypeName = "NIMLOC"
		q.TypeMeaning = "Nimrod Locator"
	case 33:
		q.TypeName = "SRV"
		q.TypeMeaning = "Server Selection"
	case 34:
		q.TypeName = "ATMA"
		q.TypeMeaning = "ATM Address"
	case 35:
		q.TypeName = "NAPTR"
		q.TypeMeaning = "Naming Authority Pointer"
	case 36:
		q.TypeName = "KX"
		q.TypeMeaning = "Key Exchanger"
	case 37:
		q.TypeName = "CERT"
		q.TypeMeaning = "CERT"
	case 38:
		q.TypeName = "A6"
		q.TypeMeaning = "A6"
	case 39:
		q.TypeName = "DNAME"
		q.TypeMeaning = "DNAME"
	case 40:
		q.TypeName = "SINK"
		q.TypeMeaning = "SINK"
	case 41:
		q.TypeName = "OPT"
		q.TypeMeaning = "OPT"
	case 42:
		q.TypeName = "APL"
		q.TypeMeaning = "APL"
	case 43:
		q.TypeName = "DS"
		q.TypeMeaning = "Delegation Signer"
	case 44:
		q.TypeName = "SSHFP"
		q.TypeMeaning = "SSH Key Fingerprint"
	case 45:
		q.TypeName = "IPSECKEY"
		q.TypeMeaning = "IPSECKEY"
	case 46:
		q.TypeName = "RRSIG"
		q.TypeMeaning = "RRSIG"
	case 47:
		q.TypeName = "NSEC"
		q.TypeMeaning = "NSEC"
	case 48:
		q.TypeName = "DNSKEY"
		q.TypeMeaning = "DNSKEY"
	case 49:
		q.TypeName = "DHCID"
		q.TypeMeaning = "DHCID"
	case 50:
		q.TypeName = "NSEC3"
		q.TypeMeaning = "NSEC3"
	case 51:
		q.TypeName = "NSEC3PARAM"
		q.TypeMeaning = "NSEC3PARAM"
	case 52:
		q.TypeName = "TLSA"
		q.TypeMeaning = "TLSA"
	case 53:
		q.TypeName = "SMIMEA"
		q.TypeMeaning = "S/MIME Certificate Association"
	case 55:
		q.TypeName = "HIP"
		q.TypeMeaning = "Host Identity Protocol"
	case 56:
		q.TypeName = "NINFO"
		q.TypeMeaning = "NINFO"
	case 57:
		q.TypeName = "RKEY"
		q.TypeMeaning = "RKEY"
	case 58:
		q.TypeName = "TALINK"
		q.TypeMeaning = "Trust Anchor LINK"
	case 59:
		q.TypeName = "CDS"
		q.TypeMeaning = "Child DS"
	case 60:
		q.TypeName = "CDNSKEY"
		q.TypeMeaning = "DNSKEY(s) The Child Wants Reflected In DS"
	case 61:
		q.TypeName = "OPENPGPKEY"
		q.TypeMeaning = "OpenPGP Key"
	case 62:
		q.TypeName = "CSYNC"
		q.TypeMeaning = "Child-To-Parent Sync"
	case 63:
		q.TypeName = "ZONEMD"
		q.TypeMeaning = "Message Digest For DNS Zone"
	case 99:
		q.TypeName = "SPF"
		q.TypeMeaning = ""
	case 100:
		q.TypeName = "UINFO"
		q.TypeMeaning = ""
	case 101:
		q.TypeName = "UID"
		q.TypeMeaning = ""
	case 102:
		q.TypeName = "GID"
		q.TypeMeaning = ""
	case 103:
		q.TypeName = "UNSPEC"
		q.TypeMeaning = ""
	case 104:
		q.TypeName = "NID"
		q.TypeMeaning = ""
	case 105:
		q.TypeName = "L32"
		q.TypeMeaning = ""
	case 106:
		q.TypeName = "L64"
		q.TypeMeaning = ""
	case 107:
		q.TypeName = "LP"
		q.TypeMeaning = ""
	case 108:
		q.TypeName = "EUI48"
		q.TypeMeaning = "An EUI-48 Address"
	case 109:
		q.TypeName = "EUI64"
		q.TypeMeaning = "An EUI-64 Address"
	case 249:
		q.TypeName = "TKEY"
		q.TypeMeaning = "Transaction Key"
	case 250:
		q.TypeName = "TSIG"
		q.TypeMeaning = "Transaction Signature"
	case 251:
		q.TypeName = "IXFR"
		q.TypeMeaning = "Incremental Transer"
	case 252:
		q.TypeName = "AXFR"
		q.TypeMeaning = "Transfer Of An Entire Zone"
	case 253:
		q.TypeName = "MAILB"
		q.TypeMeaning = "Mailbox-Related Resource Records"
	case 254:
		q.TypeName = "MAILA"
		q.TypeMeaning = "Mail Agent Resource Records"
	case 255:
		q.TypeName = "*"
		q.TypeMeaning = "A Request For Some Or All Records The Server Has Available"
	case 256:
		q.TypeName = "URI"
		q.TypeMeaning = "URI"
	case 257:
		q.TypeName = "CAA"
		q.TypeMeaning = "Certification Authority Restriction"
	case 258:
		q.TypeName = "AVC"
		q.TypeMeaning = "Application Visability And Control"
	case 259:
		q.TypeName = "DOA"
		q.TypeMeaning = "Digital Object Architecture"
	case 260:
		q.TypeName = "AMTRELAY"
		q.TypeMeaning = "Automatic Multicast Tunneling Relay"
	case 32768:
		q.TypeName = "TA"
		q.TypeMeaning = "DNSSEC Trust Authorities"
	case 32769:
		q.TypeName = "DLV"
		q.TypeMeaning = "DNSSEC Lookaside Validation"
	default:
		q.TypeName = "UNASSIGNED/PRIVATE USE/RESERVED"
		q.TypeMeaning = "UNASSIGNED/PRIVATE USE/RESERVED"
	}
}

// QueryResponse is the standard response from root-level DNS providers
type QueryResponse struct {
	StatusCode       int                     `json:"Status"`
	StatusName       string                  `json:"-"`
	StatusMessage    string                  `json:"-"`
	TC               bool                    `json:"TC"`
	RD               bool                    `json:"RD"`
	RA               bool                    `json:"RA"`
	AD               bool                    `json:"AD"`
	CD               bool                    `json:"CD"`
	Question         []QueryResponseQuestion `json:"Question"`
	Answer           []QueryResponseAnswer   `json:"Answer"`
	Additional       []interface{}           `json:"Additional"`         // Google Only
	EDNSClientSubnet string                  `json:"edns_client_subnet"` // Google Only
}

// DetermineStatusMessage will read the Status attribute and assign a message as defined by:
//     https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6
//
// Arguments:
//     None
//
// Returns:
//     None
func (q *QueryResponse) DetermineStatusMessage() {
	switch q.StatusCode {
	case 0:
		q.StatusName = "NOERROR"
		q.StatusMessage = "No Errors Reported"
	case 1:
		q.StatusName = "FORMERR"
		q.StatusMessage = "The DNS Query Is Malformed"
	case 2:
		q.StatusName = "SERVFAIL"
		q.StatusMessage = "The DNS Server Failed To Process This Request"
	case 3:
		q.StatusName = "NXDOMAIN"
		q.StatusMessage = "The Requested Domain Name Does Not Exist"
	case 4:
		q.StatusName = "NOTIMP"
		q.StatusMessage = "This is not implimented"
	case 5:
		q.StatusName = "REFUSED"
		q.StatusMessage = "The DNS Server Refused To Answer This Query"
	case 6:
		q.StatusName = "YXDOMAIN"
		q.StatusMessage = "The Requested Domain Exists, But It Should Not"
	case 7:
		q.StatusName = "YXRRSET"
		q.StatusMessage = "The Requested Resource Record Set Exists, But It Should Not"
	case 8:
		q.StatusName = "NXRRSET"
		q.StatusMessage = "The Requested Resource Record Set Does Not Exist, But It Should"
	case 9:
		q.StatusName = "NOTAUTH"
		q.StatusMessage = "Either The Server Or The Requesting User Is Not Authorized To Perform This Action"
	case 10:
		q.StatusName = "NOTZONE"
		q.StatusMessage = "The Requested Name Does Not Exist In The Requested Zone"
	case 11:
		q.StatusName = "DSOTYPEENI"
		q.StatusMessage = "DSO-TYPE: This Is Not Implimented"
	case 16:
		q.StatusName = "BADVERS/BADSIG"
		q.StatusMessage = "The Request Used A Bad OPT Version Or The TSIG Signature Failed"
	case 17:
		q.StatusName = "BADKEY"
		q.StatusMessage = "The Key Is Not Recognized"
	case 18:
		q.StatusName = "BADTIME"
		q.StatusMessage = "The Signature Is Out Of The Time Window"
	case 19:
		q.StatusName = "BADMODE"
		q.StatusMessage = "The Request Used A Bad TKEY Mode"
	case 20:
		q.StatusName = "BADNAME"
		q.StatusMessage = "The Request Used A Duplicate Key Name"
	case 21:
		q.StatusName = "BADALG"
		q.StatusMessage = "The Requested Algorithm Is Not Supported"
	case 22:
		q.StatusName = "BADTRUNC"
		q.StatusMessage = "The Requested Truncation Was Malformed"
	case 23:
		q.StatusName = "BADCOOKIE"
		q.StatusMessage = "The Server Cookie Is Bad Or Missing"
	case 65535:
		q.StatusName = "RESERVED"
		q.StatusMessage = "Reserved By Standards Action: See Provider Documentation"
	default:
		q.StatusName = "UNASSIGNED/RESERVED"
		q.StatusMessage = "Unassigned By IANA Or Reserved for Private Use (See Provider Documentation)"
	}
}

// Unmarshal is a function used to unmarshal an HTTP response body into this QueryResponse struct
//
// Arguments:
//     body (io.ReadCloser) The interface of the HTTP response body
//
// Returns:
//     (error): An error if one exists, nil otherwise
func (q *QueryResponse) Unmarshal(body io.ReadCloser) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, q)
}

// Print will print out the answers section
//
// Arguments:
//     None
//
// Returns:
//     None
func (q QueryResponse) Print() {
	fmt.Printf(
		answerStr,
		q.StatusName, q.StatusMessage,
		q.TC,
		q.RD,
		q.RA,
		q.AD,
		q.CD,
		q.EDNSClientSubnet)
	for _, i := range q.Answer {
		i.DetermineTypeNameAndMeaning()
		fmt.Printf(
			"    %s\t%d\t%s\t%s\n",
			i.Name,
			i.TTL,
			i.TypeName,
			i.Data)
	}
}

const answerStr string = `Answer:
  Status:             %s: %s
  Truncated:          %v
  RD:                 %v
  RA:                 %v
  DNSSEC Validated:   %v
  DNSSEC Disabled:    %v
  eDNS Client Subnet: %s
  Data:
`
