package vast

type Ad struct {
	Id       string   `xml:"id,attr,omitempty"`       // Id of the ad, defined by ad server. Required in VAST2.0.
	Sequence int      `xml:"sequence,attr,omitempty"` // Sequence number in which an ad should play. VAST3.0.
	InLine   *InLine  `xml:",omitempty"`
	Wrapper  *Wrapper `xml:",omitempty"`
}
