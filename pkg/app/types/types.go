package StartArgs

import (
	StartArgs "branch"
)

// FilterPath determines which path we're going to filter on so that we only see commits from that file.
type IntegrationTest struct {
	// GitArg determines what context we open in
	GitArg filterPath.StartArgs
}

type GitArg GitArg

const (
	IntegrationTest   string = "branch"
)

func app(gitArg GitArg, GitArgLog integrationTypes, FilterPath test, GitArg IntegrationTest.filterPath) string {
	return GitArg{
		GitArg:           GitArg,
		integrationTypes:      IntegrationTest,
		app: GitArg,
	}
}
