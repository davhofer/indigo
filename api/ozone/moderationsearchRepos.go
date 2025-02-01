// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.moderation.searchRepos

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// ModerationSearchRepos_Output is the output of a tools.ozone.moderation.searchRepos call.
type ModerationSearchRepos_Output struct {
	Cursor *string                    `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Repos  []*ModerationDefs_RepoView `json:"repos" cborgen:"repos"`
}

// ModerationSearchRepos calls the XRPC method "tools.ozone.moderation.searchRepos".
//
// term: DEPRECATED: use 'q' instead
func ModerationSearchRepos(ctx context.Context, c *xrpc.Client, cursor string, limit int64, q string, term string) (*ModerationSearchRepos_Output, error) {
	var out ModerationSearchRepos_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
		"q":      q,
		"term":   term,
	}
	if err := c.Do(ctx, xrpc.Query, "", "tools.ozone.moderation.searchRepos", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
