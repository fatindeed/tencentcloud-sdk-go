package cloudbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type tcbClient struct {
	*tcbConfig
	*http.Client
}

func (c *tcbClient) request(method, path string, body any) (json.RawMessage, error) {
	reqUrl := fmt.Sprintf("https://%s.api.tcloudbasegateway.com/v1/%s", c.EnvID, path)
	logrus.Debugf("Request URL: %s %s", method, reqUrl)

	var reqBody []byte
	switch body := body.(type) {
	case []byte:
		reqBody = body
	case string:
		reqBody = []byte(body)
	default:
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}
	logrus.Debugf("Request Body: %v\n", string(reqBody))

	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.GetToken())
	logrus.Debugf("Request Headers: %v\n", req.Header)

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("request failed with status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if len(respBody) == 0 {
		logrus.Warn("empty response body")
		return nil, nil
	}
	logrus.Debugf("Response Body: %s", string(respBody))

	result := &tcbResponse{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response data: %w", err)
	}
	return result.Data, nil
}
