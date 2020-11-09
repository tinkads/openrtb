package vast

type Creative struct {
	Id           string        `xml:"id,attr,omitempty"`           // ID of the creative defined by the ad server.
	Sequence     int           `xml:"sequence,attr,omitempty"`     // Sequence number in which the creative should be displayed.
	AdId         string        `xml:"AdID,attr,omitempty"`         // Id of ad associated with the creative.
	ApiFramework string        `xml:"apiFramework,attr,omitempty"` // Ad serving API used. VAST3.0.
	Linear       *Linear       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAds `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAds `xml:"NonLinearAds,omitempty"`
}
