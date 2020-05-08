package gophish

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendingProfilesService is how we access and manipulate /smtp resources.
type SendingProfilesService struct {
	Service
}

// SendingProfile is the SMTP configuration that tells Gophish how to send emails.
//
// Sending profiles support authentication and ignoring invalid SSL certificates.
type SendingProfile struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Host             string   `json:"host"`
	InterfaceType    string   `json:"interface_type"`
	FromAddress      string   `json:"from_address"`
	IgnoreCertErrors bool     `json:"ignore_cert_errors"`
	ModifiedDate     string   `json:"modified_date"`
	Headers          []Header `json:"headers"`
}

// Header is a <key,value> pair that we send along with the SendingProfile.
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ListSendingProfiles retrieves a list of the sending profiles created by the
// authenticated user.
func (ss *SendingProfilesService) ListSendingProfiles() ([]SendingProfile, error) {
	resp, err := ss.MakeRequest("GET", "/api/smtp", nil)
	if err != nil {
		return nil, err
	}

	var profiles []SendingProfile
	if err := json.NewDecoder(resp.Body).Decode(&profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}

// GetSendingProfile returns a sending profile given an ID, returning a 404
// error if no sending profile with the provided ID is found.
func (ss *SendingProfilesService) GetSendingProfile(id int) (*SendingProfile, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/smtp/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var profile SendingProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// CreateSendingProfile creates a sending profile.
func (ss *SendingProfilesService) CreateSendingProfile(sp *SendingProfile) (*SendingProfile, error) {
	resp, err := ss.MakeRequest("POST", "/api/smtp", sp)
	if err != nil {
		return nil, err
	}

	var profile SendingProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateSendingProfile modifies an existing sending profile.
func (ss *SendingProfilesService) UpdateSendingProfile(sp *SendingProfile) (*SendingProfile, error) {
	resp, err := ss.MakeRequest(
		"PUT",
		fmt.Sprintf("/api/smtp/%d", sp.ID),
		sp,
	)
	if err != nil {
		return nil, err
	}

	var profile SendingProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// DeleteSendingProfile deletes a sending profile by ID.
func (ss *SendingProfilesService) DeleteSendingProfile(id int) (bool, error) {
	resp, err := ss.MakeRequest(
		"DELETE",
		fmt.Sprintf("/api/smtp/%d", id),
		nil,
	)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}
