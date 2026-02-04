package cloudbase

import (
	"encoding/json"
	"fmt"
)

type databaseClient struct {
	*tcbClient
}

func (c *databaseClient) List(req *ModelListRequest) (*ModelListResult, error) {
	path := fmt.Sprintf("model/prod/%s/list", req.Name)
	data, err := c.request("POST", path, req)
	if err != nil {
		return nil, err
	}
	result := &ModelListResult{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response data: %w", err)
	}
	return result, nil
}

func (c *databaseClient) Update(req *ModelUpdateRequest) (*ModeUpdateResult, error) {
	path := fmt.Sprintf("model/prod/%s/update", req.Name)
	data, err := c.request("PUT", path, req)
	if err != nil {
		return nil, err
	}
	result := &ModeUpdateResult{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response data: %w", err)
	}
	return result, nil
}
