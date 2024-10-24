package profile
type ProfileService struct {
	ProfileRepository
}

func New(repo ProfileRepository) (*ProfileService, error) {
	return &ProfileService{ProfileRepository: repo}, nil
}
