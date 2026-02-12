package gitpanda_fetcher

import (
	"net/http"
	"strings"

	"github.com/cockroachdb/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

const titleSeparator = " · "

// Client represents GitLab URL parser
type Client struct {
	baseURL        string
	client         *gitlab.Client
	isDebugLogging bool
}

// ClientParams represents parameters of NewClient
type ClientParams struct {
	APIEndpoint    string
	BaseURL        string
	PrivateToken   string
	UserAgent      string
	IsDebugLogging bool
	HTTPClient     *http.Client
}

// NewClient create new Client instance
func NewClient(params *ClientParams) (*Client, error) {
	p := &Client{
		baseURL:        params.BaseURL,
		isDebugLogging: params.IsDebugLogging,
	}

	options := []gitlab.ClientOptionFunc{gitlab.WithBaseURL(params.APIEndpoint)}
	if params.HTTPClient != nil {
		options = append(options, gitlab.WithHTTPClient(params.HTTPClient))
	}
	client, err := gitlab.NewClient(params.PrivateToken, options...)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	p.client = client

	if params.UserAgent == "" {
		p.client.UserAgent = "gitpanda_fetcher (+https://github.com/sue445/gitpanda_fetcher)"
	} else {
		p.client.UserAgent = params.UserAgent
	}

	return p, nil
}

// FetchURL fetch GitLab url
func (p *Client) FetchURL(url string) (*Page, error) {
	if !strings.HasPrefix(url, p.baseURL) {
		return nil, nil
	}

	pos := len(p.baseURL)
	if !strings.HasSuffix(p.baseURL, "/") {
		pos++
	}
	path := url[pos:]

	fetchers := []fetcher{
		&epicFetcher{},
		&snippetFetcher{},
		&issueFetcher{},
		&mergeRequestFetcher{},
		&jobFetcher{},
		&pipelineFetcher{},
		&blobFetcher{},
		&commitFetcher{},
		&projectSnippetFetcher{},
		&projectFetcher{},
		&userOrGroupFetcher{},
	}

	for _, fetcher := range fetchers {
		page, err := fetcher.fetchPath(path, p.client, p.isDebugLogging)

		if err != nil {
			return nil, errors.WithStack(err)
		}

		if page != nil {
			return page, nil
		}
	}

	return nil, nil
}
