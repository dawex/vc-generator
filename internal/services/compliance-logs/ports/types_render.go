package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

// Render : Pre-processing before a response is marshalled and sent across the wire
func (a ComplianceLog) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind : just a post-process after a decode..
func (na *ComplianceLogIn) Bind(r *http.Request) error {
	return nil
}

func ConvertComplianceLogListRenders(complianceLogList []*ComplianceLog) []render.Renderer {
	list := []render.Renderer{}
	for _, complianceLog := range complianceLogList {
		list = append(list, complianceLog)
	}
	return list
}
