package protocols

// Generic
type GenericRequest struct {
	Command string `json:"command"`
}

// Custom Request
type CreateOrUpdateSiteRequest struct {
	GenericRequest
	Site
}
type GetSiteRequest struct {
	GenericRequest
	Params string `json:"params"`
}

type CreateOrUpdateFileRequest struct {
	GenericRequest
	File
}
type GetFileRequest struct {
	GenericRequest
	File
}
