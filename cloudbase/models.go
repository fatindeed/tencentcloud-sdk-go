package cloudbase

import (
	"encoding/json"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type tcbConfig struct {
	*common.Credential
	EnvID string
}

type tcbResponse struct {
	RequestID string          `json:"requestId"`
	Data      json.RawMessage `json:"data"`
}

type ModelFilter struct {
	Where map[string]any `json:"where,omitempty"`
}

// https://docs.cloudbase.net/http-api/model/get-records
type ModelListRequest struct {
	Name       string         `json:"-"`
	Filter     *ModelFilter   `json:"filter,omitempty"`
	Select     map[string]int `json:"select,omitempty"`
	PageSize   int            `json:"pageSize,omitempty"`
	PageNumber int            `json:"pageNumber,omitempty"`
	GetCount   bool           `json:"getCount,omitempty"`
	OrderBy    map[string]int `json:"orderBy,omitempty"`
}

type ModelListResult struct {
	Records []json.RawMessage `json:"records"`
	Total   int               `json:"total"`
}

// https://docs.cloudbase.net/http-api/model/update-item
type ModelUpdateRequest struct {
	Name   string         `json:"-"`
	Filter *ModelFilter   `json:"filter"`
	Data   map[string]any `json:"data"`
}

type ModeUpdateResult struct {
	Count int `json:"count"`
}
