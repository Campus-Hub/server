package service

const (
	// File Types
	Video    = "Video"
	Document = "Document"
	Picture  = "Picture"
)

type (
	ResourceService struct {
		Name   string
		Format string
		Type   string
	}
)

func (service ResourceService) Upload() {

}
