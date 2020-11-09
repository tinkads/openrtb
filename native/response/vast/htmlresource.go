package vast

type HtmlResource struct {
	Html         TrimmedData `xml:",cdata"`
	IsXmlEncoded bool        `xml:"xmlEncoded,attr,omitempty"`
}
