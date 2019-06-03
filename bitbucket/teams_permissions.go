package bitbucket

type TeamPermissions struct {
	Pagination

	Values []*TeamPermission `json:"values,omitempty"`
}

type TeamPermission struct {
	Permission *string `json:"permission,omitempty"`
	Type       *string `json:"type,omitempty"`
	User       *User   `json:"user,omitempty"`
	Team       *Team   `json:"team,omitempty"`
}

type TeamRepoPermissions struct {
	Pagination

	Values []*TeamRepoPermission `json:"values,omitempty"`
}

type TeamRepoPermission struct {
	Permission *string     `json:"permission,omitempty"`
	Type       *string     `json:"type,omitempty"`
	User       *User       `json:"user,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
}

// ListPermissions returns each team permission a user on the team has.
//
// Permissions returned are effective permissions — if a user is a member of multiple groups with distinct roles,
// only the highest level is returned.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/permissions#get
func (t *TeamsService) ListPermissions(teamUsername string) (*TeamPermissions, *Response, error) {
	result := new(TeamPermissions)
	urlStr := t.client.requestUrl("/teams/%s/permissions", teamUsername)
	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// ListRepositoryPermissions returns each repository permission for all of a team’s repositories.
//
// If the username URL parameter refers to a user account instead of a team account,
// an object containing the repository permissions of all the username's repositories will be returned.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/permissions/repositories#get
func (t *TeamsService) ListRepositoryPermissions(teamUsername string, opts *FilterSortOpts) (*TeamRepoPermissions, *Response, error) {
	result := new(TeamRepoPermissions)
	urlStr := t.client.requestUrl("/teams/%s/permissions/repositories", teamUsername)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}

// GetRepositoryPermissions returns each repository permission of a given repository.
//
// If the username URL parameter refers to a user account instead of a team account,
// an object containing the repository permissions of the username's repository will be returned.
//
// Bitbucket API docs: https://developer.atlassian.com/bitbucket/api/2/reference/resource/teams/%7Busername%7D/permissions/repositories/%7Brepo_slug%7D#get
func (t *TeamsService) GetRepositoryPermissions(teamUsername, repoSlug string, opts *FilterSortOpts) (*TeamRepoPermissions, *Response, error) {
	result := new(TeamRepoPermissions)
	urlStr := t.client.requestUrl("/teams/%s/permissions/repositories/%s", teamUsername, repoSlug)
	urlStr, addOptErr := addOptions(urlStr, opts)
	if addOptErr != nil {
		return nil, nil, addOptErr
	}

	response, err := t.client.execute("GET", urlStr, result, nil)

	return result, response, err
}
