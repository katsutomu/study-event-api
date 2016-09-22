package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// SearchKeywordsPath computes a request path to the search action of keywords.
func SearchKeywordsPath(text string) string {
	return fmt.Sprintf("/search/%v", text)
}

// 複数の勉強会サイトをキーワード検索してレスポンスを返す
func (c *Client) SearchKeywords(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSearchKeywordsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSearchKeywordsRequest create the request corresponding to the search action endpoint of the keywords resource.
func (c *Client) NewSearchKeywordsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
