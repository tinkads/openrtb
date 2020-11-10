package vast

type StaticResource struct {
	MimeType string      `xml:"creativeType,attr"` // Required.
	Uri      TrimmedData `xml:",cdata"`
}
