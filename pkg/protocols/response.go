package protocols

// Generic
type GenericResponse struct {
	Command string `json:"command"`
	Status  int    `json:"status"`
}

// Custom Response
type CreateOrUpdateSiteResponse struct {
	GenericResponse
	Site
}
type GetSiteResponse struct {
	GenericResponse
	Sites []Site
}

type CreateOrUpdateFileResponse struct {
	GenericResponse
	File
}
type GetFileResponse struct {
	GenericResponse
	Files []File
}
