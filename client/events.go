package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// KeepEventEventsPath computes a request path to the keep event action of events.
func KeepEventEventsPath(eventID int) string {
	return fmt.Sprintf("/api/v2/events/%v/keep", eventID)
}

// イベントのお気に入り操作
func (c *Client) KeepEventEvents(ctx context.Context, path string, isKeep bool) (*http.Response, error) {
	req, err := c.NewKeepEventEventsRequest(ctx, path, isKeep)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewKeepEventEventsRequest create the request corresponding to the keep event action endpoint of the events resource.
func (c *Client) NewKeepEventEventsRequest(ctx context.Context, path string, isKeep bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp10 := strconv.FormatBool(isKeep)
	values.Set("isKeep", tmp10)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// ListEventsPath computes a request path to the list action of events.
func ListEventsPath(id string) string {
	return fmt.Sprintf("/api/v2/events/genre/%v", id)
}

// ListEventsPath2 computes a request path to the list action of events.
func ListEventsPath2() string {
	return fmt.Sprintf("/api/v2/events/new")
}

// ListEventsPath3 computes a request path to the list action of events.
func ListEventsPath3() string {
	return fmt.Sprintf("/api/v2/events/keep")
}

// ListEventsPath4 computes a request path to the list action of events.
func ListEventsPath4() string {
	return fmt.Sprintf("/api/v2/events/nokeep")
}

// ListEventsPath5 computes a request path to the list action of events.
func ListEventsPath5() string {
	return fmt.Sprintf("/api/v2/events/popular")
}

// ListEventsPath6 computes a request path to the list action of events.
func ListEventsPath6() string {
	return fmt.Sprintf("/api/v2/events/recommend")
}

// イベント情報取得
func (c *Client) ListEvents(ctx context.Context, path string, page *int, q *string, sort *string) (*http.Response, error) {
	req, err := c.NewListEventsRequest(ctx, path, page, q, sort)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEventsRequest create the request corresponding to the list action endpoint of the events resource.
func (c *Client) NewListEventsRequest(ctx context.Context, path string, page *int, q *string, sort *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if page != nil {
		tmp11 := strconv.Itoa(*page)
		values.Set("page", tmp11)
	}
	if q != nil {
		values.Set("q", *q)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}
