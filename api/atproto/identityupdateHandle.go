// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.identity.updateHandle

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// IdentityUpdateHandle_Input is the input argument to a com.atproto.identity.updateHandle call.
type IdentityUpdateHandle_Input struct {
	// handle: The new handle.
	Handle string `json:"handle" cborgen:"handle"`
}

// IdentityUpdateHandle calls the XRPC method "com.atproto.identity.updateHandle".
func IdentityUpdateHandle(ctx context.Context, c *xrpc.Client, input *IdentityUpdateHandle_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.identity.updateHandle", nil, input, nil); err != nil {
		return err
	}

	return nil
}
