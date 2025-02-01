// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.identity.resolveHandle

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// IdentityResolveHandle_Output is the output of a com.atproto.identity.resolveHandle call.
type IdentityResolveHandle_Output struct {
	Did string `json:"did" cborgen:"did"`
}

// IdentityResolveHandle calls the XRPC method "com.atproto.identity.resolveHandle".
//
// handle: The handle to resolve.
func IdentityResolveHandle(ctx context.Context, c *xrpc.Client, handle string) (*IdentityResolveHandle_Output, error) {
	var out IdentityResolveHandle_Output

	params := map[string]interface{}{
		"handle": handle,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.identity.resolveHandle", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
