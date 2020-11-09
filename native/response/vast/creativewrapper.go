package vast

type CreativeWrapper struct {
	Id           string               `xml:"id,attr,omitempty"`
	Sequence     int                  `xml:"sequence,attr,omitempty"`
	AdId         string               `xml:"AdID,attr,omitempty"`
	Linear       *LinearWrapper       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAdsWrapper `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAdsWrapper `xml:"NonLinearAds,omitempty"`
}
