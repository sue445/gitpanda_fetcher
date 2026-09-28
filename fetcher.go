package fetcher

import (
	"regexp"

	"gitlab.com/gitlab-org/api/client-go/v3"
)

// reProjectName represents regular expression matching to a project name
const reProjectName = `^([^/]+(?:/[^/]+)+)`

type fetcher interface {
	fetchPath(path string, client *gitlab.Client, isDebugLogging bool) (*Page, error)
}

func sanitizeProjectName(projectName string) string {
	return regexp.MustCompile(`(/-)$`).ReplaceAllString(projectName, "")
}
