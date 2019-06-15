package routing

// PasteRequest is the data from a incoming paste request
type PasteRequest struct {
	ExpirationTime string
	PasteContent   string
	CustomURL      string
}

// PasteResponse is the response for a paste
type PasteResponse struct {
	URL string
}
