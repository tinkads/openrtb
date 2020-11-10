package vast

type CompanionAds struct {
	Required   Mode         `xml:"required,attr,omitempty"` // VAST3.0.
	Companions []*Companion `xml:"Companion,omitempty"`
}
