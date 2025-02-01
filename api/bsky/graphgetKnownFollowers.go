// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.getKnownFollowers

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// GraphGetKnownFollowers_Output is the output of a app.bsky.graph.getKnownFollowers call.
type GraphGetKnownFollowers_Output struct {
	Cursor    *string                  `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Followers []*ActorDefs_ProfileView `json:"followers" cborgen:"followers"`
	Subject   *ActorDefs_ProfileView   `json:"subject" cborgen:"subject"`
}

// GraphGetKnownFollowers calls the XRPC method "app.bsky.graph.getKnownFollowers".
func GraphGetKnownFollowers(ctx context.Context, c *xrpc.Client, actor string, cursor string, limit int64) (*GraphGetKnownFollowers_Output, error) {
	var out GraphGetKnownFollowers_Output

	params := map[string]interface{}{
		"actor":  actor,
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.graph.getKnownFollowers", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
