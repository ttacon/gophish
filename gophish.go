package gophish

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func NewClient(host, token string) *Client {
	service := Service{
		Host:  host,
		Token: token,
	}
	return &Client{
		SendingProfiles: SendingProfilesService{service},
		Templates:       TemplatesService{service},
		LandingPages:    LandingPagesService{service},
		Groups:          GroupsService{service},
		Campaigns:       CampaignsService{service},
	}
}

type Client struct {
	SendingProfiles SendingProfilesService
	Templates       TemplatesService
	LandingPages    LandingPagesService
	Groups          GroupsService
	Campaigns       CampaignsService
}

type Service struct {
	Host  string
	Token string
}

func (s Service) MakeRequest(method, path string, payload interface{}) (*http.Response, error) {
	var r io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		r = bytes.NewBuffer(data)
	}

	path = s.Host + path
	req, err := http.NewRequest(method, path, r)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", s.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
