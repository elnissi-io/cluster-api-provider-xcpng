package xcpng

import (
	"net/http"
	"time"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NewClient returns a new Client configured to communicate with the XCP-ng server.
func NewClient(cfg Config) Client {
	return &client{
		serverURL: cfg.ServerURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		username: cfg.Username,
		password: cfg.Password,
	}
}

type client struct {
	serverURL  string
	httpClient *http.Client
	username   string
	password   string
}

// CreateVM implements the creation of a virtual machine on the XCP-ng server.
func (c *client) CreateVM(ctx context.Context, vmSpec VMSpec) (string, error) {
	requestData := map[string]interface{}{
		"method": "vm.create",
		"params": []interface{}{vmSpec},
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	vmID, ok := result["result"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	return vmID, nil
}
