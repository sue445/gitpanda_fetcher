# gitpanda_fetcher
Module for parsing GitLab URLs and retrieving information via GitLab API

[![test](https://github.com/sue445/gitpanda_fetcher/actions/workflows/test.yml/badge.svg)](https://github.com/sue445/gitpanda_fetcher/actions/workflows/test.yml)

## Usage
```go
package main

import (
	"github.com/sue445/gitpanda_fetcher"
)

func main() {
	client, err := fetcher.NewClient(&fetcher.URLParserParams{
		APIEndpoint:    "https://gitlab.example.com/api/v4",
		BaseURL:        "https://gitlab.example.com",
		PrivateToken:   "xxxxxxxxxx",
		IsDebugLogging: true,
	})

	page, err := client.FetchURL("https://gitlab.example.com/user/repo/-/merge_requests/123")
}
```

## Requirements
* GitLab API v4

## Supported URL format
* User URL
    * e.g. `${GITLAB_BASE_URL}/:username`
* Group URL
    * e.g. `${GITLAB_BASE_URL}/:groupname`
* Project URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame`
* Issue URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/issues/:iid`
* MergeRequest URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/merge_requests/:iid`
* Job URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/-/jobs/:id`
* Pipeline URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/pipelines/:id`
* Blob URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/blob/:sha1/:filename`
* Commit URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/commit/:sha1`
* Project snippet URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/snippets/:id`
* Snippet URL
    * e.g. `${GITLAB_BASE_URL}/snippets/:id`
* Epic URL
    * e.g. `${GITLAB_BASE_URL}/groups/:groupname/-/epics/:iid`
* Work item URL
    * e.g. `${GITLAB_BASE_URL}/:namespace/:reponame/-/work_items/:iid`
