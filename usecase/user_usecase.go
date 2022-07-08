package usecase

import "api/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(a domain.UserRepository) domain.UserUsecase {
	return &userUsecase{a}
}

func (u *userUsecase) FindUser(id string) (bool, *domain.User) {
	user := u.userRepo.FindUser(id)
	if user == nil {
		return false, nil
	} else {
		return true, user
	}
}

func (u *userUsecase) AddUser(id string) bool {
	c := u.userRepo.AddUser(id)
	return c
}

func (u *userUsecase) DeleteUser(id string) bool {
	c := u.userRepo.DeleteUser(id)
	return c
}

func (u *userUsecase) UpdateUser(id string, name string) bool {
	c := u.userRepo.UpdateUser(id, name)
	return c
}

func (u *userUsecase) GetUser() []domain.User {
	users := u.userRepo.GetUser()
	return users
}
