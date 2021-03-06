package usecase

import (
	"ddd-sample/domain/model"
	"ddd-sample/domain/repository"
	"ddd-sample/domain/service"
)

type UserUsecase interface {
	Index() (*[]model.User, error)
	Find(id uint) (*model.User, error)
	Create(name, email, password string) (*model.User, error)
	Update(id uint, name, email, password string) error
	Delete(id uint) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: ur}
}

func (uu *userUsecase) Index() (*[]model.User, error) {
	users, err := uu.userRepository.All()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *userUsecase) Find(id uint) (*model.User, error) {
	u, err := uu.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uu *userUsecase) Create(name, email, password string) (*model.User, error) {
	u, err := model.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	us := service.NewUser(u.GetEmail())
	ok, err := us.Exists(uu.userRepository, u)
	if ok {
		return nil, err
	}
	u, err = uu.userRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uu *userUsecase) Update(id uint, name, email, password string) error {
	u, err := model.NewUser(name, email, password)
	if err != nil {
		return err
	}
	u.SetID(id)
	us := service.NewUser(u.GetEmail())
	ok, err := us.Exists(uu.userRepository, u)
	if ok {
		return err
	}
	if err := uu.userRepository.Update(u, id); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) Delete(id uint) error {
	if err := uu.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
