package common

type ContentType string

const (
	JSON      = ContentType("application/json")
	XML       = ContentType("application/xml")
	PLAIN     = ContentType("text/plain")
	HTML      = ContentType("text/html")
	FILE_FORM = ContentType("multipart/form-data")
	FILE_SOCKET = ContentType("multipart/socket-data")
	FORM      = ContentType("application/x-www-form-urlencoded")
)