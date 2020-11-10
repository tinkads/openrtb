package vast

type NonLinearAds struct {
	Trackings  []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears []*NonLinear `xml:"NonLinear,omitempty"` // Required, at least one item.
}
