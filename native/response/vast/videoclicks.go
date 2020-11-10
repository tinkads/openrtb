package vast

type VideoClicks struct {
	ClickThroughs  []*VideoClick `xml:"ClickThrough,omitempty"`
	ClickTrackings []*VideoClick `xml:"ClickTracking,omitempty"`
	CustomClicks   []*VideoClick `xml:"CustomClick,omitempty"`
}
