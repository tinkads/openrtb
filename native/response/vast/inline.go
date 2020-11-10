package vast

type InLine struct {
	AdSystem    *AdSystem     `xml:"AdSystem"`           // Required.
	AdTitle     *string       `xml:"AdTitle"`            // Required.
	Impressions []*Impression `xml:"Impression"`         // Required.
	Creatives   []*Creative   `xml:"Creatives>Creative"` // Creatives node is required.
	Description string        `xml:"Description,omitempty"`
	Advertiser  string        `xml:"Advertiser,omitempty"` // VAST3.0.
	SurveyUrl   string        `xml:"Survey,omitempty"`
	Errors      []string      `xml:"Error,omitempty"`
	Pricing     *Pricing      `xml:"Pricing,omitempty"` // VAST3.0.
	Extensions  []*Extension  `xml:"Extensions>Extension,omitempty"`
}
