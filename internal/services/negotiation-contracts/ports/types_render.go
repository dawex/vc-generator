package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

// Render : Pre-processing before a response is marshalled and sent across the wire
func (a NegotiationContract) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind : just a post-process after a decode..
func (na *NegotiationContract) Bind(r *http.Request) error {
	return nil
}

func ConvertNegotiationContractListRenders(listIn []*NegotiationContract) []render.Renderer {
	list := []render.Renderer{}
	for _, elem := range listIn {
		list = append(list, elem)
	}
	return list
}
