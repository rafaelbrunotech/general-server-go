package getusers

type GetUsersQueryInput struct{}

type GetUsersQuery struct{}

func NewGetUsersQuery(input GetUsersQueryInput) (*GetUsersQuery, error) {
	return &GetUsersQuery{}, nil
}
