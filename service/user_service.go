package service

import (
	"errors"
	"go-hm2/models"
)

type UserService struct {
	users  map[int]*models.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) *models.User {
	user := &models.User{
		ID:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}

	s.users[s.nextID] = user
	s.nextID++

	return user
}

func (s *UserService) GetById(id int) (*models.User, error) {
	user, exists := s.users[id]

	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	user, exists := s.users[id]

	if !exists {
		return nil, errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	_, exists := s.users[id]

	if !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}

func (s *UserService) GetAll() []*models.User {
	userList := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		userList = append(userList, user)
	}

	return userList
}
