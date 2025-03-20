package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

// Render : Pre-processing before a response is marshalled and sent across the wire
func (a PublicKey) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render : Pre-processing before a response is marshalled and sent across the wire
func (a VcSigned) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ConvertVcSignedListRenders(eventList []*VcSigned) []render.Renderer {
	list := []render.Renderer{}
	for _, event := range eventList {
		list = append(list, event)
	}
	return list
}
