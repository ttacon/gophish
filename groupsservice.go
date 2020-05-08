package gophish

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Group is a set of recipients that allows us to manage recipients in bulk
// for campaigns.
type Group struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Targets      []Target `json:"targets"`
	ModifiedDate string   `json:"modified_date"`
}

// Target is a specific Gophish target.
type Target struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
}

// GroupsService is how we access and manipulate /groups.
type GroupsService struct {
	Service
}

// ListGroups returns a list of groups.
func (ss *GroupsService) ListGroups() ([]Group, error) {
	resp, err := ss.MakeRequest("GET", "/api/groups", nil)
	if err != nil {
		return nil, err
	}

	var groups []Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// GetGroup retrieves a group given an ID.
func (ss *GroupsService) GetGroup(id int) (*Group, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/groups/%d", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, err
	}
	return &group, nil
}

// CreateGroup creates a new group.
func (ss *GroupsService) CreateGroup(sp *Group) (*Group, error) {
	resp, err := ss.MakeRequest("POST", "/api/groups", sp)
	if err != nil {
		return nil, err
	}

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, err
	}
	return &group, nil
}

// UpdateGroup updates the given group.
func (ss *GroupsService) UpdateGroup(sp *Group) (*Group, error) {
	resp, err := ss.MakeRequest(
		"PUT",
		fmt.Sprintf("/api/groups/%d", sp.ID),
		sp,
	)
	if err != nil {
		return nil, err
	}

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, err
	}
	return &group, nil
}

// DeleteGroup deletes a group given an ID.
func (ss *GroupsService) DeleteGroup(id int) (bool, error) {
	resp, err := ss.MakeRequest(
		"DELETE",
		fmt.Sprintf("/api/groups/%d", id),
		nil,
	)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}

// ImportGroup imports a group from a CSV - this is not currently supported.
func (ss *GroupsService) ImportGroup(imp ImportGroupRequest) (*Group, error) {
	return nil, errors.New("importing groups isn't currently supported")
}

// ImportGroupRequest is the payload for importing a new group.
type ImportGroupRequest struct {
	File string
}

// ListGroupSummaries returns a list of group summaries.
func (ss *GroupsService) ListGroupSummaries() ([]GroupSummary, error) {
	resp, err := ss.MakeRequest("GET", "/api/groups/summary", nil)
	if err != nil {
		return nil, err
	}

	var groups []GroupSummary
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// GetGroupSummary retrieves a group summary given a group ID.
func (ss *GroupsService) GetGroupSummary(id int) (*GroupSummary, error) {
	resp, err := ss.MakeRequest(
		"GET",
		fmt.Sprintf("/api/groups/%d/summary", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var group GroupSummary
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		return nil, err
	}
	return &group, nil
}

// GroupSummary is a top-level summary of a group (i.e. how many targets it has
// and when it was last modified).
type GroupSummary struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	NumTargets   []int  `json:"num_targets"`
	ModifiedDate string `json:"modified_date"`
}
