package vast

type CompanionAdsWrapper struct {
	Required   Mode                `xml:"required,attr,omitempty"` // VAST3.0.
	Companions []*CompanionWrapper `xml:"Companion,omitempty"`
}
