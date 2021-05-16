package notion

import (
	"context"
	"fmt"
)

type PageService service

// PropertyValueEnum ...
type PropertyValueEnum int

const (
	Test PropertyValueEnum = iota

)

// Page represents the property values of a single Notion page
type Page struct {
	Object string `json:"object"`
	ID string `json:"id"`
	CreatedTime string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	Archived bool `json:"archived"`
	// Properties Property `json:"properties"`
}

// Property ...
type Property struct {
	Key string `json:"key"`
	Value PropertyValue `json:"value"`
}

// PropertyValue represents the identifier, type, and value of a page property
type PropertyValue  struct {
	ID string
	Type string
}


// DatabaseParent ...
type DatabaseParent struct {
	Type string
	DatabaseID string
}

// PageParent ...
type PageParent struct {
	Type string
	PageID string
}

// WorkspaceParent represents a page with a workspace parent is a top-level page within a Notion workspace.
// The parent property is an object containing the following keys:
type WorkspaceParent struct {
	Type string
}

// Create will
func (p *PageService) Create(ctx context.Context) error {
	return nil
}

// Retrieve returns a Page object using the specified ID
func (p *PageService) Retrieve(ctx context.Context, pageID string) (*Page, error) {
	req, err := p.client.NewRequest("GET", fmt.Sprintf("pages/%s", pageID), nil)
	if err != nil {
		return nil, err
	}

	var page *Page
	_, err = p.client.Do(ctx, req, &page)
	if err != nil {
		return nil, err
	}

	return page, nil
}


// helpers

func (p PropertyValueEnum) validate() string {
	return [...]string{"test"}[p]
}
