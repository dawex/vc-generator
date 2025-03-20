package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

// Render : Pre-processing before a response is marshalled and sent across the wire
func (a Event) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind : just a post-process after a decode..
func (na *EventIn) Bind(r *http.Request) error {
	return nil
}

func ConvertEventListRenders(eventList []*Event) []render.Renderer {
	list := []render.Renderer{}
	for _, event := range eventList {
		list = append(list, event)
	}
	return list
}
