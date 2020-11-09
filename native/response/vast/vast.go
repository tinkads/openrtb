package vast

import "encoding/xml"

type Vast struct {
	XMLName xml.Name `xml:"VAST"`
	Version Version  `xml:"version,attr"`    // Required.
	Ads     []*Ad    `xml:"Ad"`              // Ad can be empty in VAST2.0.
	Errors  []string `xml:"Error,omitempty"` // One or more URI's, likely tracking pixels, to request in case of no ad.
}
