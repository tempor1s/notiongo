package notion

import (
	"context"
	"fmt"
)

type PageService service

// Page represents the property values of a single Notion page
type Page struct {
	Object         string        `json:"object"`
	ID             string        `json:"id"`
	CreatedTime    string        `json:"created_time"`
	LastEditedTime string        `json:"last_edited_time"`
	Archived       bool          `json:archived"`
	Properties     PropertyValue `json:"properties"`
}

// PropertyValue represents the identifier, type, and value of a page property
type PropertyValue struct {
	ID   string
	Type string
}

// Retrieve returns a Notion Page object using the specified ID
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

// Update mutates an existing Notion Page object using the specified ID
//
// example:
// '{
// 	"properties": {
// 	  "In stock": { "checkbox": true }
// 	}
//   }'
//
func (p *PageService) Update(ctx context.Context, pageID string, updates Property) (*Page, error) {
	return nil, nil
}

// Create creates a new Notion Page object in the specified database or as a child of an existing page
func (p *PageService) Create(ctx context.Context) (*Page, error) {
	return nil, nil
}

// helpers
