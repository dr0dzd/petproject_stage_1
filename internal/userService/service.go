package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) CreateUser(u User) (User, error) {
	return s.repo.CreateUser(u)
}

func (s *UserService) UpdateUserByID(id uint, u User) (User, error) {
	return s.repo.UpdateUserByID(id, u)
}

func (s *UserService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
