package bitbucket

import "time"

// CommitComments represent a collection of a commit's comments.
type CommitComments struct {
	PaginationInfo

	Values []*CommitComment `json:"values,omitempty"`
}

// CommitComment represents a commit comment.
type CommitComment struct {
	Links     *CCLinks          `json:"links,omitempty"`
	Deleted   *bool             `json:"deleted,omitempty"`
	Content   *BitbucketContent `json:"content,omitempty"`
	CreatedOn *time.Time        `json:"created_on,omitempty"`
	User      *User             `json:"user,omitempty"`
	Commit    *Commit           `json:"commit,omitempty"`
	UpdatedOn *time.Time        `json:"updated_on,omitempty"`
	Type      *string           `json:"type,omitempty"`
	ID        *int64            `json:"id,omitempty"`
}

// CCLinks represents the "links" object in a Bitbucket commit comment.
type CCLinks struct {
	Self *BitbucketLink `json:"self,omitempty"`
	HTML *BitbucketLink `json:"html,omitempty"`
}

// CommitCommentRequest represents a new commit comment.
type CommitCommentRequest struct {
	Content       *BitbucketContent `json:"content,omitempty"`
	ParentComment struct {
		Id *int64 `json:"id,omitempty"`
	} `json:"parent,omitempty"`
}

// ListComments returns the commit's comments.
//
// This includes both global and inline comments.
// The default sorting is oldest to newest and can be overridden with the sort query parameter.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/comments#get
func (c *CommitService) ListComments(owner, repoSlug, sha string, opts ...interface{}) (*CommitComments, *Response, error) {
	results := new(CommitComments)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commit/%s/comments", owner, repoSlug, sha)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("GET", urlStr, results, nil)

	return results, response, err
}

// CreateComment creates new comment on the specified commit.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/comments#post
func (c *CommitService) CreateComment(owner, repoSlug, sha string, co *CommitCommentRequest, opts ...interface{}) (*CommitComment, *Response, error) {
	results := new(CommitComment)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commit/%s/comments", owner, repoSlug, sha)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("POST", urlStr, results, co)

	return results, response, err
}

// GetComment returns the specified commit comment.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/comments/%7Bcomment_id%7D#get
func (c *CommitService) GetComment(owner, repoSlug, sha string, cID int64, opts ...interface{}) (*CommitComment, *Response, error) {
	results := new(CommitComment)
	urlStr := c.client.requestUrl("/repositories/%s/%s/commit/%s/comments/%v", owner, repoSlug, sha, cID)
	urlStr, addOptErr := addOptions(urlStr, opts...)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := c.client.execute("POST", urlStr, results, nil)

	return results, response, err
}
