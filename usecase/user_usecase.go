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

func (u *userUsecase) AddUser(user domain.User) bool {
	c := u.userRepo.AddUser(user)
	return c
}

func (u *userUsecase) DeleteUser(id string) bool {
	c := u.userRepo.DeleteUser(id)
	return c
}

func (u *userUsecase) UpdateUser(user domain.User, id string) bool {
	c := u.userRepo.UpdateUser(user, id)
	return c
}

func (u *userUsecase) GetUser() []domain.User {
	users := u.userRepo.GetUser()
	return users
}
