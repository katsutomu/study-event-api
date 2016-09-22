//************************************************************************//
// API "study-event": Application Contexts
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
)

// SearchKeywordsContext provides the keywords search action context.
type SearchKeywordsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Text string
}

// NewSearchKeywordsContext parses the incoming request URL and body, performs validations and creates the
// context used by the keywords controller search action.
func NewSearchKeywordsContext(ctx context.Context, service *goa.Service) (*SearchKeywordsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SearchKeywordsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramText := req.Params["text"]
	if len(paramText) > 0 {
		rawText := paramText[0]
		rctx.Text = rawText
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SearchKeywordsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "applicateion/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
