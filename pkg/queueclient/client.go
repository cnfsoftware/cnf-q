package queueclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	URL    string
	Token  string
	client *http.Client
}

func NewClient(url, token string) *Client {
	return &Client{
		URL:    url,
		Token:  token,
		client: http.DefaultClient,
	}
}

func (c *Client) Push(queue string, body []byte) error {
	url := c.URL + "/queue/" + queue + "/push"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	c.addAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

func (c *Client) Pop(queue string) ([]byte, error) {
	url := c.URL + "/queue/" + queue + "/pop"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.addAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (c *Client) Peek(queue string) ([]byte, error) {
	url := c.URL + "/queue/" + queue + "/peek"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.addAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

type ListQueue struct {
	Queues []string `json:"queues,omitempty"`
	Error  *string  `json:"error,omitempty"`
}

func (c *Client) ListQueues() ([]string, error) {
	url := c.URL + "/queues"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	c.addAuthHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	queues := ListQueue{}
	err = json.NewDecoder(resp.Body).Decode(&queues)
	if err != nil {
		return nil, err
	}
	if queues.Error != nil {
		return nil, errors.New(*queues.Error)
	}

	return queues.Queues, nil
}

func (c *Client) addAuthHeader(req *http.Request) {
	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}
}
