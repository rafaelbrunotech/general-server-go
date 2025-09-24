package getusers

type GetUsersQueryInput struct{}

type GetUsersQuery struct{}

func NewQuery(input GetUsersQueryInput) (*GetUsersQuery, error) {
	return &GetUsersQuery{}, nil
}
