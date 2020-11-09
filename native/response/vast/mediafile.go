package vast

type MediaFile struct {
	Id                        string      `xml:"id,attr,omitempty"`
	Delivery                  Delivery    `xml:"delivery,attr"`             // Required.
	MimeType                  string      `xml:"type,attr"`                 // Required.
	Codec                     string      `xml:"codec,attr,omitempty"`      // VAST3.0.
	Bitrate                   *int        `xml:"bitrate,attr,omitempty"`    // In Kbps; absent if MaxBitrate and MinBitrate are present.
	MinBitrate                *int        `xml:"minBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
	MaxBitrate                *int        `xml:"maxBitrate,attr,omitempty"` // In Kbps; absent if Bitrate is present. VAST3.0.
	Width                     int         `xml:"width,attr"`                // Required.
	Height                    int         `xml:"height,attr"`               // Required.
	IsScalable                bool        `xml:"scalable,attr,omitempty"`
	ShouldMaintainAspectRatio bool        `xml:"maintainAspectRatio,attr,omitempty"`
	ApiFramework              string      `xml:"apiFramework,attr,omitempty"` // API used to interact with the MediaFile.
	Uri                       TrimmedData `xml:",cdata"`
}
