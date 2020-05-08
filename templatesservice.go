package gophish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Template is the content of the emails that are sent to targets. They can be
// imported from an existing email, or created from scratch.
//
// Additionally, templates can contain tracking images so that gophish knows
// when the user opens the email.
type Template struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Subject      string       `json:"subject"`
	Text         string       `json:"text"`
	HTML         string       `json:"html"`
	ModifiedDate string       `json:"modified_date"`
	Attachments  []Attachment `json:"attachments"`
}

// Attachment is an attachment in a given Template.
type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}

// TemplatesService is how we access and manipulate /templates.
type TemplatesService struct {
	Service
}

// ListTemplates returns a list of templates.
func (ss *TemplatesService) ListTemplates() ([]Template, error) {
	resp, err := ss.MakeRequest("GET", "/api/templates", nil)
	if err != nil {
		return nil, err
	}

	var templates []Template
	if err := json.NewDecoder(resp.Body).Decode(&templates); err != nil {
		return nil, err
	}
	return templates, nil
}

// GetTemplate retrieves a template given an ID.
func (ss *TemplatesService) GetTemplate(id int) (*Template, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/templates/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var profile Template
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// CreateTemplate creates a new template.
func (ss *TemplatesService) CreateTemplate(sp *Template) (*Template, error) {
	resp, err := ss.MakeRequest("POST", "/api/templates", sp)
	if err != nil {
		return nil, err
	}

	var profile Template
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateTemplate updates the given template.
func (ss *TemplatesService) UpdateTemplate(sp *Template) (*Template, error) {
	resp, err := ss.MakeRequest(
		"PUT",
		fmt.Sprintf("/api/templates/%d", sp.ID),
		sp,
	)
	if err != nil {
		return nil, err
	}

	var template Template
	if err := json.NewDecoder(resp.Body).Decode(&template); err != nil {
		return nil, err
	}
	return &template, nil
}

// DeleteTemplate deletes a template given an ID.
func (ss *TemplatesService) DeleteTemplate(id int) (bool, error) {
	resp, err := ss.MakeRequest(
		"DELETE",
		fmt.Sprintf("/api/templates/%d", id),
		nil,
	)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}

// ImportTemplate imports a template from a given email.
//
// This makes it easy to repurpose legitimate emails for your phishing
// assessments.
//
// This endpoint expects the raw email content in RFC 2045 format, including
// the original headers. Usually, this is found using the "Show Original"
// feature of email clients.
func (ss *TemplatesService) ImportTemplate(imp ImportRequest) (*Template, error) {
	resp, err := ss.MakeRequest(
		"POST",
		"/api/import/email",
		imp,
	)
	if err != nil {
		return nil, err
	}

	var template Template
	if err := json.NewDecoder(resp.Body).Decode(&template); err != nil {
		return nil, err
	}
	return &template, nil
}

// ImportRequest is the payload for importing a new template from an email.
type ImportRequest struct {
	ConvertLinks bool   `json:"convert_links"`
	Content      string `json:"content"`
}
