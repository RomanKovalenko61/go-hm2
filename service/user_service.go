package service

import (
	"fmt"
	"go-hm2/models"
	"sync"
)

type UserService struct {
	users  map[int]*models.User
	nextID int
	mutex  sync.RWMutex
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

func (s *UserService) Create(req *models.CreateUserRequest) *models.User {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, s.userNotFoundError(id)
	}

	return user, nil
}

func (s *UserService) Update(id int, req *models.UpdateUserRequest) (*models.User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, s.userNotFoundError(id)
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	return user, nil
}

func (s *UserService) Delete(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, exists := s.users[id]
	if !exists {
		return s.userNotFoundError(id)
	}
	delete(s.users, id)

	return nil
}

func (s *UserService) GetAll() []*models.User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	userList := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		userList = append(userList, user)
	}

	return userList
}

func (s *UserService) userExists(id int) (*models.User, bool) {
	user, exists := s.users[id]
	return user, exists
}

func (s *UserService) userNotFoundError(id int) error {
	return fmt.Errorf("user with ID %d not found", id)
}
