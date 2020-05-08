package gophish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LandingPage is the HTML content returned when targets click on the links
// in Gophish emails.
type LandingPage struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	HTML               string `json:"html"`
	CaptureCredentials bool   `json:"capture_credentials"`
	CapturePasswords   bool   `json:"capture_passwords"`
	ModifiedDate       string `json:"modified_date"`
	RedirectURL        string `json:"redirect_url"`
}

// LandingPagesService is how we access and manipulate /pages resources.
type LandingPagesService struct {
	Service
}

// ListLandingPages returns a list of landing pages.
func (ss *LandingPagesService) ListLandingPages() ([]LandingPage, error) {
	resp, err := ss.MakeRequest("GET", "/api/pages", nil)
	if err != nil {
		return nil, err
	}

	var landingpages []LandingPage
	if err := json.NewDecoder(resp.Body).Decode(&landingpages); err != nil {
		return nil, err
	}
	return landingpages, nil
}

// GetLandingPage returns a landing page given an ID.
func (ss *LandingPagesService) GetLandingPage(id int) (*LandingPage, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/pages/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var profile LandingPage
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// CreateLandingPage creates a new landing page.
func (ss *LandingPagesService) CreateLandingPage(sp *LandingPage) (*LandingPage, error) {
	resp, err := ss.MakeRequest("POST", "/api/pages", sp)
	if err != nil {
		return nil, err
	}

	var profile LandingPage
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateLandingPage updates the given landing page.
func (ss *LandingPagesService) UpdateLandingPage(sp *LandingPage) (*LandingPage, error) {
	resp, err := ss.MakeRequest(
		"PUT",
		fmt.Sprintf("/api/pages/%d", sp.ID),
		sp,
	)
	if err != nil {
		return nil, err
	}

	var landingpage LandingPage
	if err := json.NewDecoder(resp.Body).Decode(&landingpage); err != nil {
		return nil, err
	}
	return &landingpage, nil
}

// DeleteLandingPage deletes a landing page given an ID.
func (ss *LandingPagesService) DeleteLandingPage(id int) (bool, error) {
	resp, err := ss.MakeRequest(
		"DELETE",
		fmt.Sprintf("/api/pages/%d", id),
		nil,
	)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}

// ImportSite fetches a URL to be later imported as a landing page.
//
// This endpoint simply fetches and returns the HTML from a provided URL. If
// include_resources is false (recommended), a <base> tag is added so that
// relative links in the HTML resolve from the original URL.
//
// Additionally, if the HTML contains form elements, this endpoint adds another
// input, __original_url, that points to the original URL. This makes it
// possible to replay captured credentials later.
func (ss *LandingPagesService) ImportSite(imp ImportSiteRequest) (*LandingPage, error) {
	resp, err := ss.MakeRequest(
		"POST",
		"/api/import/site",
		imp,
	)
	if err != nil {
		return nil, err
	}

	var landingpage LandingPage
	if err := json.NewDecoder(resp.Body).Decode(&landingpage); err != nil {
		return nil, err
	}
	return &landingpage, nil
}

type ImportSiteRequest struct {
	IncludeResources bool   `json:"include_resources"`
	URL              string `json:"url"`
}
