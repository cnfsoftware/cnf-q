package queueclient

import (
	"bytes"
	"io"
	"net/http"
)

type Client struct {
	URL    string
	client *http.Client
}

func NewClient(url string) *Client {
	return &Client{
		URL:    url,
		client: http.DefaultClient,
	}
}

func (c *Client) Push(queue string, body []byte) error {
	url := c.URL + "/" + queue + "/push"
	_, err := c.client.Post(url, "application/json", bytes.NewBuffer(body))
	return err
}

func (c *Client) Pop(queue string) ([]byte, error) {
	url := c.URL + "/" + queue + "/pop"
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return body, nil
}

func (c *Client) Peek(queue string) ([]byte, error) {
	url := c.URL + "/" + queue + "/peek"
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return body, nil
}
