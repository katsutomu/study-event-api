//************************************************************************//
// API "study-event": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=study-event-api/design
// --out=$(GOPATH)/src/study-event-api
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// KeywordsController is the controller interface for the Keywords actions.
type KeywordsController interface {
	goa.Muxer
	Search(*SearchKeywordsContext) error
}

// MountKeywordsController "mounts" a Keywords resource controller on the given service.
func MountKeywordsController(service *goa.Service, ctrl KeywordsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSearchKeywordsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Search(rctx)
	}
	service.Mux.Handle("GET", "/search/:text", ctrl.MuxHandler("Search", h, nil))
	service.LogInfo("mount", "ctrl", "Keywords", "action", "Search", "route", "GET /search/:text")
}
