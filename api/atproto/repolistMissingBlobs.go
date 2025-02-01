// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.repo.listMissingBlobs

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// RepoListMissingBlobs_Output is the output of a com.atproto.repo.listMissingBlobs call.
type RepoListMissingBlobs_Output struct {
	Blobs  []*RepoListMissingBlobs_RecordBlob `json:"blobs" cborgen:"blobs"`
	Cursor *string                            `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
}

// RepoListMissingBlobs_RecordBlob is a "recordBlob" in the com.atproto.repo.listMissingBlobs schema.
type RepoListMissingBlobs_RecordBlob struct {
	Cid       string `json:"cid" cborgen:"cid"`
	RecordUri string `json:"recordUri" cborgen:"recordUri"`
}

// RepoListMissingBlobs calls the XRPC method "com.atproto.repo.listMissingBlobs".
func RepoListMissingBlobs(ctx context.Context, c *xrpc.Client, cursor string, limit int64) (*RepoListMissingBlobs_Output, error) {
	var out RepoListMissingBlobs_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.repo.listMissingBlobs", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
