package util

// PasteRequest is the data from a incoming paste request
type PasteRequest struct {
	ExpirationTime string
	PasteContent   string
	CustomURL      string
}

// PasteResponse is the response for a paste
type PasteResponse struct {
	URL string `json: url`
}

// PasteData is the type of the data in the DB
type PasteData struct {
	ExpLength   string
	CreatedTime string
	PastePath   string
	Data        string
}

// DataResponse is the pasted data
type DataResponse struct {
	Data string
}
