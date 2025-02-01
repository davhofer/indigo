// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.notification.putPreferences

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// NotificationPutPreferences_Input is the input argument to a app.bsky.notification.putPreferences call.
type NotificationPutPreferences_Input struct {
	Priority bool `json:"priority" cborgen:"priority"`
}

// NotificationPutPreferences calls the XRPC method "app.bsky.notification.putPreferences".
func NotificationPutPreferences(ctx context.Context, c *xrpc.Client, input *NotificationPutPreferences_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "app.bsky.notification.putPreferences", nil, input, nil); err != nil {
		return err
	}

	return nil
}
