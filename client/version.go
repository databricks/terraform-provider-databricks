package client

type APIVersion string

const (
	// APIVersion is the version of the RESTful API of DataBricks
	API2  APIVersion = "2.0"
	API12 APIVersion = "1.2"

	// SdkVersion is the version of this library
	SdkVersion = "0.0.1"
)
