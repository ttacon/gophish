package gophish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Campaign is our phishing campaign against a given group of recipients.
type Campaign struct {
	ID            int              `json:"id"`
	Name          string           `json:"name"`
	CreatedDate   string           `json:"created_date"`
	LaunchDate    string           `json:"launch_date"`
	SendByDate    string           `json:"send_by_date"`
	CompletedDate string           `json:"completed_date"`
	Template      Template         `json:"template"`
	Page          LandingPage      `json:"page"`
	Status        string           `json:"status"`
	Stats         CampaignStats    `json:"stats"`
	Results       []CampaignResult `json:"result"`
	Groups        []Group          `json:"groups"`
	Timeline      []CampaignEvent  `json:"timeline"`
	SMTP          SendingProfile   `json:"smtp"`
	URL           string           `json:"url"`
}

// CampaignStats is the top-level stats for the results of a given campaign.
type CampaignStats struct {
	Total         int `json:"total"`
	Sent          int `json:"sent"`
	Opened        int `json:"opened"`
	Clicked       int `json:"clicked"`
	SubmittedData int `json:"submitted_data"`
	EmailReported int `json:"email_reported"`
}

// CampaignEvent is an event in the lifetime of a campaign. This includes, and
// is not limited to:
//
//  - Sending emails
//  - Recipients opening emails
//  - Recipients clicking on phishing links
//  - Recipients entering credentials into phishing sites
//  - Recipients reporting phishing emails
type CampaignEvent struct {
	Email   string `json:"email"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// CampaignResult is a specific result for a given recipient in a given
// campaign.
type CampaignResult struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Position  string  `json:"position"`
	Status    string  `json:"status"`
	IP        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	SendDate  string  `json:"send_date"`
	Reported  bool    `json:"reported"`
}

// CampaignsService is how we access and manipulate /campaigns.
type CampaignsService struct {
	Service
}

// ListCampaigns retrieves a list of campaigns.
func (ss *CampaignsService) ListCampaigns() ([]Campaign, error) {
	resp, err := ss.MakeRequest("GET", "/api/campaigns", nil)
	if err != nil {
		return nil, err
	}

	var campaigns []Campaign
	if err := json.NewDecoder(resp.Body).Decode(&campaigns); err != nil {
		return nil, err
	}
	return campaigns, nil
}

// GetCampaign retrieves a campaign given an ID.
func (ss *CampaignsService) GetCampaign(id int) (*Campaign, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/campaigns/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

// GetCampaignResults retrieves the results for a campaign given and ID.
func (ss *CampaignsService) GetCampaignResults(id int) (*Campaign, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/campaigns/%d/results", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

// GetCampaignSummary retrieves the summer for a campaign given and ID.
func (ss *CampaignsService) GetCampaignSummary(id int) (*Campaign, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/campaigns/%d/summary", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

// CreateCampaign creates a new campaign.
func (ss *CampaignsService) CreateCampaign(sp *Campaign) (*Campaign, error) {
	resp, err := ss.MakeRequest("POST", "/api/campaigns", sp)
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

// UpdateCampaign updates the given campaign.
func (ss *CampaignsService) UpdateCampaign(sp *Campaign) (*Campaign, error) {
	resp, err := ss.MakeRequest(
		"PUT",
		fmt.Sprintf("/api/campaigns/%d", sp.ID),
		sp,
	)
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	if err := json.NewDecoder(resp.Body).Decode(&campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}

// DeleteCampaign deletes a campaign given an ID.
func (ss *CampaignsService) DeleteCampaign(id int) (bool, error) {
	resp, err := ss.MakeRequest(
		"DELETE",
		fmt.Sprintf("/api/campaigns/%d", id),
		nil,
	)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}

// CompleteCampaign marks a campaign as completed given an ID.
func (ss *CampaignsService) CompleteCampaign(id int) (bool, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/campaigns/%d", id),
		nil,
	)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}
