// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.feed.describeFeedGenerator

import (
	"context"

	"github.com/davhofer/indigo/xrpc"
)

// FeedDescribeFeedGenerator_Feed is a "feed" in the app.bsky.feed.describeFeedGenerator schema.
type FeedDescribeFeedGenerator_Feed struct {
	Uri string `json:"uri" cborgen:"uri"`
}

// FeedDescribeFeedGenerator_Links is a "links" in the app.bsky.feed.describeFeedGenerator schema.
type FeedDescribeFeedGenerator_Links struct {
	PrivacyPolicy  *string `json:"privacyPolicy,omitempty" cborgen:"privacyPolicy,omitempty"`
	TermsOfService *string `json:"termsOfService,omitempty" cborgen:"termsOfService,omitempty"`
}

// FeedDescribeFeedGenerator_Output is the output of a app.bsky.feed.describeFeedGenerator call.
type FeedDescribeFeedGenerator_Output struct {
	Did   string                            `json:"did" cborgen:"did"`
	Feeds []*FeedDescribeFeedGenerator_Feed `json:"feeds" cborgen:"feeds"`
	Links *FeedDescribeFeedGenerator_Links  `json:"links,omitempty" cborgen:"links,omitempty"`
}

// FeedDescribeFeedGenerator calls the XRPC method "app.bsky.feed.describeFeedGenerator".
func FeedDescribeFeedGenerator(ctx context.Context, c *xrpc.Client) (*FeedDescribeFeedGenerator_Output, error) {
	var out FeedDescribeFeedGenerator_Output
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.describeFeedGenerator", nil, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
