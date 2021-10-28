package auth

type Authorization interface {
	CreateUser() error
}

type AuthService struct {
	storage Authorization
}

func NewAuthService(storage *Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) CreateUser() error {
	err := s.storage.CreateUser()
	return err
}
