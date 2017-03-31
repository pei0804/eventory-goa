package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// AccountCreateUsersPath computes a request path to the account create action of users.
func AccountCreateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/new")
}

// 正規ユーザーの作成
func (c *Client) AccountCreateUsers(ctx context.Context, path string, clientVersion *string, email *string, identifier *string, platform *string) (*http.Response, error) {
	req, err := c.NewAccountCreateUsersRequest(ctx, path, clientVersion, email, identifier, platform)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAccountCreateUsersRequest create the request corresponding to the account create action endpoint of the users resource.
func (c *Client) NewAccountCreateUsersRequest(ctx context.Context, path string, clientVersion *string, email *string, identifier *string, platform *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if clientVersion != nil {
		values.Set("client_version", *clientVersion)
	}
	if email != nil {
		values.Set("email", *email)
	}
	if identifier != nil {
		values.Set("identifier", *identifier)
	}
	if platform != nil {
		values.Set("platform", *platform)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.KeySigner != nil {
		c.KeySigner.Sign(req)
	}
	return req, nil
}

// TmpAccountCreateUsersPath computes a request path to the tmp account create action of users.
func TmpAccountCreateUsersPath() string {
	return fmt.Sprintf("/api/v2/users/tmp")
}

// 一時ユーザーの作成
func (c *Client) TmpAccountCreateUsers(ctx context.Context, path string, clientVersion *string, identifier *string, platform *string) (*http.Response, error) {
	req, err := c.NewTmpAccountCreateUsersRequest(ctx, path, clientVersion, identifier, platform)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTmpAccountCreateUsersRequest create the request corresponding to the tmp account create action endpoint of the users resource.
func (c *Client) NewTmpAccountCreateUsersRequest(ctx context.Context, path string, clientVersion *string, identifier *string, platform *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if clientVersion != nil {
		values.Set("client_version", *clientVersion)
	}
	if identifier != nil {
		values.Set("identifier", *identifier)
	}
	if platform != nil {
		values.Set("platform", *platform)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
