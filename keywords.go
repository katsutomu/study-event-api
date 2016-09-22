package main

import (
	"github.com/goadesign/goa"
	"study-event-api/app"
)

// KeywordsController implements the keywords resource.
type KeywordsController struct {
	*goa.Controller
}

// NewKeywordsController creates a keywords controller.
func NewKeywordsController(service *goa.Service) *KeywordsController {
	return &KeywordsController{Controller: service.NewController("KeywordsController")}
}

// Search runs the search action.
func (c *KeywordsController) Search(ctx *app.SearchKeywordsContext) error {
	// KeywordsController_Search: start_implement

	// Put your logic here

	// KeywordsController_Search: end_implement
	return nil
}
