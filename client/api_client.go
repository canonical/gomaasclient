package client

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type OAuthToken struct {
	ConsumerKey string
	TokenKey    string
	TokenSecret string
}

func NewOAuthToken(apiKey string) (*OAuthToken, error) {
	parts := strings.Split(apiKey, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid API key format: %q; expected '<consumer>:<token>:<secret>'", apiKey)
	}

	return &OAuthToken{
		ConsumerKey: parts[0],
		TokenKey:    parts[1],
		TokenSecret: parts[2],
	}, nil
}

type APIClient struct {
	token   *OAuthToken
	baseURL *url.URL
	client  *http.Client
}

func NewAPIClient(apiURL, apiVersion, apiKey string, client *http.Client) (*APIClient, error) {
	token, err := NewOAuthToken(apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}

	baseURL, err := url.Parse(fmt.Sprintf("%sapi/%s/", ensureTrailingSlash(apiURL), apiVersion))
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if client == nil {
		client = http.DefaultClient
	}

	return &APIClient{
		token:   token,
		baseURL: baseURL,
		client:  client,
	}, nil
}

func (c *APIClient) SubClient(resource string) *APIClient {
	newURL := c.baseURL.ResolveReference(&url.URL{Path: resource})
	newURL.Path = ensureTrailingSlash(newURL.Path)

	return &APIClient{
		token:   c.token,
		baseURL: newURL,
		client:  c.client,
	}
}

func (c *APIClient) SubClientWithoutSlash(resource string) *APIClient {
	newURL := c.baseURL.ResolveReference(&url.URL{Path: resource})

	return &APIClient{
		token:   c.token,
		baseURL: newURL,
		client:  c.client,
	}
}

func (c *APIClient) setOAuthHeader(req *http.Request) error {
	nonce, err := generateNonce()
	if err != nil {
		return err
	}

	authParams := map[string]string{
		"realm":                  "MAAS API",
		"oauth_consumer_key":     c.token.ConsumerKey,
		"oauth_token":            c.token.TokenKey,
		"oauth_signature_method": "PLAINTEXT",
		"oauth_signature":        "&" + c.token.TokenSecret,
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_nonce":            nonce,
		"oauth_version":          "1.0",
	}

	var headerParts []string
	for k, v := range authParams {
		headerParts = append(headerParts, fmt.Sprintf(`%s="%s"`, k, url.QueryEscape(v)))
	}

	req.Header.Set("Authorization", "OAuth "+strings.Join(headerParts, ", "))

	return nil
}

func (c *APIClient) doRequest(ctx context.Context, method, operation string, params url.Values, contentType string, body io.Reader, handler func([]byte) error) error {
	queryURL := *c.baseURL

	if params == nil {
		params = make(url.Values)
	}

	if operation != "" {
		params.Set("op", operation)
	}

	queryURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), body)
	if err != nil {
		return fmt.Errorf("failed to create %s request: %w", method, err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	err = c.setOAuthHeader(req)
	if err != nil {
		return fmt.Errorf("failed to set authorization for %s: %w", method, err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do %s request: %w", method, err)
	}
	defer resp.Body.Close() //nolint:errcheck // ignore error

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := io.ReadAll(resp.Body) //nolint:errcheck // can be empty

		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(data))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read %s response: %w", method, err)
	}

	return handler(data)
}

func (c *APIClient) Get(ctx context.Context, operation string, params url.Values, handler func([]byte) error) error {
	return c.doRequest(ctx, http.MethodGet, operation, params, "", nil, handler)
}

func (c *APIClient) Post(ctx context.Context, operation string, params url.Values, handler func([]byte) error) error {
	contentType := "application/x-www-form-urlencoded"
	body := strings.NewReader(params.Encode())

	return c.doRequest(ctx, http.MethodPost, operation, nil, contentType, body, handler)
}

func (c *APIClient) PostFiles(ctx context.Context, operation string, params url.Values, files map[string][]byte, handler func([]byte) error) error {
	contentType, body, err := createMultipartFormData(params, files)
	if err != nil {
		return fmt.Errorf("failed to create multipart form: %w", err)
	}

	return c.doRequest(ctx, http.MethodPost, operation, nil, contentType, body, handler)
}

func (c *APIClient) Put(ctx context.Context, params url.Values, handler func([]byte) error) error {
	contentType := "application/x-www-form-urlencoded"
	body := strings.NewReader(params.Encode())

	return c.doRequest(ctx, http.MethodPut, "", nil, contentType, body, handler)
}

func (c *APIClient) PutFiles(ctx context.Context, params url.Values, files map[string][]byte, handler func([]byte) error) error {
	contentType, body, err := createMultipartFormData(params, files)
	if err != nil {
		return fmt.Errorf("failed to create multipart form: %w", err)
	}

	return c.doRequest(ctx, http.MethodPut, "", nil, contentType, body, handler)
}

func (c *APIClient) Delete(ctx context.Context) error {
	return c.doRequest(ctx, http.MethodDelete, "", nil, "", nil, func(b []byte) error { return nil })
}

func createMultipartFormData(params url.Values, files map[string][]byte) (string, io.Reader, error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	for key, values := range params {
		for _, value := range values {
			fw, err := writer.CreateFormField(key)
			if err != nil {
				return "", nil, fmt.Errorf("failed to create form field: %w", err)
			}

			if _, err := io.Copy(fw, bytes.NewBufferString(value)); err != nil {
				return "", nil, fmt.Errorf("failed to write form field: %w", err)
			}
		}
	}

	for fileName, fileContent := range files {
		fw, err := writer.CreateFormFile(fileName, fileName)
		if err != nil {
			return "", nil, fmt.Errorf("failed to create form file: %w", err)
		}

		if _, err := io.Copy(fw, bytes.NewBuffer(fileContent)); err != nil {
			return "", nil, fmt.Errorf("failed to write file data: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return "", nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	return writer.FormDataContentType(), buf, nil
}

func ensureTrailingSlash(urlStr string) string {
	if !strings.HasSuffix(urlStr, "/") {
		return urlStr + "/"
	}

	return urlStr
}

func generateNonce() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	return fmt.Sprintf("%16x", bytes), nil
}
