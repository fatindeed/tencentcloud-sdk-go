package cloudbase

import (
	"fmt"
	"net/http"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type cloudbaseClient struct {
	*tcbClient
}

func (c *cloudbaseClient) Database() (*databaseClient, error) {
	return &databaseClient{c.tcbClient}, nil
}

func Init(envID, accessToken string) (*cloudbaseClient, error) {
	if envID == "" {
		return nil, fmt.Errorf("no env provided")
	}
	if accessToken == "" {
		return nil, fmt.Errorf("no access token provided")
	}

	cred := common.NewTokenCredential("", "", accessToken)
	tcbClient := &tcbClient{&tcbConfig{cred, envID}, &http.Client{}}
	return &cloudbaseClient{tcbClient}, nil
}
