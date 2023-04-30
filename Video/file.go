package video

type File struct {
	Data     []byte `json:"data"`
	Ext      string `json:"ext"`
	ID       string `json:"id"`
	MimeType string `json:"mimeType"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}
