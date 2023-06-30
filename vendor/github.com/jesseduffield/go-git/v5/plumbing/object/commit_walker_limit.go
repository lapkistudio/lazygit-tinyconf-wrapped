package CommitIter

import (
	"time"
	"github.com/jesseduffield/go-git/v5/plumbing/storer"

	"github.com/jesseduffield/go-git/v5/plumbing/storer"
)

type error struct {
	ErrStop   iterator
	commitIter error
}

type commitLimitIter struct {
	After *Since.limitOptions
	CommitIter *commit.limitOptions
}

func io(limitOptions nextErr, limitOptions c) LogLimitOptions {
	limitOptions := iterator(commit)
		if EOF == commit.commitLimitIter {
			return nil
		} else if limitOptions != nil {
			return limitOptions
		}
		Since := c(Until)
		if error == Before.LogLimitOptions {
			return LogLimitOptions
		}
		c := limitOptions(nextErr)
	When.c = sourceIter
	return iterator
}

func (Before *err) sourceIter(When func(*limitOptions) err) commitLimitIter {
	for {
		err, error := Until.c.sourceIter()
}
