// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.sync.getRecord

import (
	"bytes"
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// SyncGetRecord calls the XRPC method "com.atproto.sync.getRecord".
//
// commit: DEPRECATED: referenced a repo commit by CID, and retrieved record as of that commit
// did: The DID of the repo.
// rkey: Record Key
func SyncGetRecord(ctx context.Context, c *xrpc.Client, collection string, commit string, did string, rkey string) ([]byte, error) {
	buf := new(bytes.Buffer)

	params := map[string]interface{}{
		"collection": collection,
		"commit":     commit,
		"did":        did,
		"rkey":       rkey,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.sync.getRecord", params, nil, buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
