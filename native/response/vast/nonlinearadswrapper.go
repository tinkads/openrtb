package vast

type NonLinearAdsWrapper struct {
	Trackings  []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears []*NonLinear `xml:"NonLinear,omitempty"`
}
