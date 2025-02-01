package rules

import (
	"strings"

	appbsky "github.com/davhofer/indigo/api/bsky"
	"github.com/davhofer/indigo/automod"
)

func MisleadingLinkUnicodeReversalPostRule(c *automod.RecordContext, post *appbsky.FeedPost) error {

	if !strings.Contains(post.Text, "\u202E") {
		return nil
	}

	c.AddRecordFlag("clickjack-unicode-reversed")
	return nil
}
