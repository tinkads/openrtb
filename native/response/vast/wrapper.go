package vast

type Wrapper struct {
	AdSystem     *AdSystem          `xml:"AdSystem"`     // Required.
	VastAdTagUri string             `xml:"VASTAdTagURI"` // Required.
	Impressions  []*Impression      `xml:"Impression"`   // Required.
	Errors       []string           `xml:"Error,omitempty"`
	Creatives    []*CreativeWrapper `xml:"Creatives>Creative"` // Required.
	Extensions   []*Extension       `xml:"Extensions>Extension,omitempty"`
}
